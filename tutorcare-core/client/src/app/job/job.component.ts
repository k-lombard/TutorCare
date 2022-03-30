import {Component, OnInit, ChangeDetectionStrategy, Inject} from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { MatDialog, MatDialogConfig, MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { MatFormFieldControl } from '@angular/material/form-field';
import { ActivatedRoute, NavigationEnd, Router } from '@angular/router';
import { select, Store } from '@ngrx/store';
import { DefaultGlobalConfig, ToastrService } from 'ngx-toastr';
import { BehaviorSubject, Observable, Subject, Subscription } from 'rxjs';
import { map, switchMap } from 'rxjs/operators';
import { getCurrUser, isLoggedIn } from 'src/app/auth/auth.selectors';
import { Application } from 'src/app/models/application.model';
import { Post } from 'src/app/models/post.model';
import { User } from 'src/app/models/user.model';
import { AppState } from 'src/app/reducers';
import { GeolocationPositionWithUser } from '../models/geolocationposition.model';
import { JobService } from './job.service';
import { NgModule, CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';
import { DateValidator } from '../find-jobs/date.validator';
import { PostCode } from '../models/postcode.model';

@Component({
  selector: 'job-component',
  templateUrl: './job.component.html',
  styleUrls: ['./job.component.scss']
})
export class JobComponent implements OnInit {

    public acceptSubject: Subject<boolean> = new BehaviorSubject<boolean>(false);
    public acceptActive = this.acceptSubject.asObservable();
    verifyCode: BehaviorSubject<number> = new BehaviorSubject<number>(0);
    active: BehaviorSubject<boolean> = new BehaviorSubject<boolean>(false);
    verified: BehaviorSubject<boolean> = new BehaviorSubject<boolean>(false);
    selectedValue: string | undefined
    userCategory: string = ""
    user!: User
    userId!: string
    posts!: Post[]
    currPost!: Post
    menuVisible: boolean
    mainCol: boolean
    locs!: GeolocationPositionWithUser[]
    mySubscription!: any
    postId!: number
    userType!: string
    start!: Date
    editable: boolean = true
    codeForm: FormGroup;
    private routeSub: Subscription;

    constructor(private router: Router, public dialog: MatDialog, private route: ActivatedRoute, private store: Store<AppState>, private toastr: ToastrService, private jobService: JobService, private fb: FormBuilder) {}

    async ngOnInit() {
        this.routeSub = this.route.params.subscribe(params => {
            this.postId = parseInt(params['id'])
        });
        this.store
        .pipe(
            select(getCurrUser)
        ).subscribe(data =>  {
            this.user = data
            this.userId = this.user.user_id || ""
            this.userType = this.user.user_category
        })
        if (this.postId && (!this.posts || this.posts.length === 0)) {
            await this.jobService.getPostByIdPromise(this.postId).then((resp: PostCode) => {
                this.currPost = resp
                console.log(resp)
                this.start = new Date(this.currPost.start_date)
                this.start.setHours(parseInt(this.currPost.start_time.split(':')[0]))
            })
        }
        this.codeForm = this.fb.group({
            code: new FormControl('', Validators.compose([
            Validators.maxLength(25),
            Validators.pattern('^[0-9]*$'),
            Validators.required
            ])),
        })
    
        if (this.userId === this.currPost?.user_id || this.userId === this.currPost?.caregiver_id) {
            if(this.currPost?.title) {
                var now = new Date()
                // if (now.getDate() === this.start.getDate() && now.getHours() - this.start.getHours() < 1 && now.getHours() - this.start.getHours() > 0) {
                    this.active.next(true);
                    this.jobService.getVerificationCode(this.currPost.post_id).subscribe((out: PostCode) => {
                        this.verifyCode.next(out?.code)
                        this.verified.next(out?.verified)
                    })
                // }
            }
        } else {
            this.router.navigate(['/home'])
        }
    }

    ngOnDestroy() {
      this.routeSub.unsubscribe();
    }

    onCodeSubmit() {
        var code: number = parseInt(this.codeForm.get('code').value)
        this.jobService.updatePostCode(this.currPost.post_id, code).subscribe(resp => {
            this.verified.next(true)
        })
    }

    onPosterCompleted() {
        this.jobService.updateJobPostPoster(this.currPost.post_id, this.currPost).subscribe((resp: Post) => {
            this.currPost.poster_completed = true;
        })
    }

    onCaregiverCompleted() {
        this.jobService.updateJobPostCaregiver(this.currPost.post_id, this.currPost).subscribe((resp: Post) => {
            this.currPost.caregiver_completed = true;
        })
    }
    

    onFindCareClick() {
        this.router.navigate(['/find-care'])
    }

    setPost(post: Post) {
      this.currPost = post
    }

    onDeleteClick(post: Post) {
      if(window.confirm("Are you sure you want to delete the job post: " + post.title + "?"))
      this.jobService.deletePost(post.post_id).subscribe(id => {
        this.currPost = undefined
        this.toastr.success("Post successfully deleted with id: "+ id, "Success", {closeButton: true, timeOut: 5000, progressBar: true});
        var postsCopy: Post[] = []
        for (let pos of this.posts) {
          if (pos.post_id != post.post_id) {
            postsCopy.push(pos)
          }
        }
        this.posts = postsCopy
      })
    }

    onEditClick() {
      this.editable = false
    }

    setSelected(i: number) {
      this.jobService.setSelectedIdx(i)
    }

    getSelected() {
      return this.jobService.getSelected()
    }

    back() {
      this.currPost= undefined
    }

    backToMenu() {
      this.mainCol = false
      this.menuVisible = true
    }

    onSaveClick() {
      this.editable = true
      this.jobService.editJobPost(this.currPost.user_id, this.currPost.post_id, this.currPost.title, this.currPost.tags, this.currPost.care_description, this.currPost.start_date, this.currPost.start_time, this.currPost.end_date, this.currPost.end_time, this.currPost.care_type).subscribe(updatedPost => {
        console.log(updatedPost)
      })
    }
}

interface FilterOption {
  value: string;
  viewValue: string;
}

interface Tag {
  display: string,
  value: string
}
