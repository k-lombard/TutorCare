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
import { GeolocationPositionWithUser } from '../../models/geolocationposition.model';
import { HistoryService } from './history.service';
import { NgModule, CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';
import { DateValidator } from '../../find-jobs/date.validator';
import { PostCode } from '../../models/postcode.model';

@Component({
  selector: 'history-component',
  templateUrl: './history.component.html',
  styleUrls: ['./history.component.scss']
})
export class HistoryComponent implements OnInit {

    public acceptSubject: Subject<boolean> = new BehaviorSubject<boolean>(false);
    public acceptActive = this.acceptSubject.asObservable();
    verifyCode: BehaviorSubject<number> = new BehaviorSubject<number>(0);
    active: BehaviorSubject<boolean> = new BehaviorSubject<boolean>(false);
    verified: BehaviorSubject<boolean> = new BehaviorSubject<boolean>(false);
    selectedValue: string | undefined
    userCategory: string = ""
    user!: User
    userId!: string
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
    posts!: Post[]
    private routeSub: Subscription;

    constructor(private router: Router, public dialog: MatDialog, private route: ActivatedRoute, private store: Store<AppState>, private toastr: ToastrService, private historyService: HistoryService, private fb: FormBuilder) {}

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
        await this.historyService.getPostsByUserIdCompleted(this.userId).then((resp) => {
            this.posts = resp.posts
            console.log(this.posts)
        })
    }

    ngOnDestroy() {
      this.routeSub.unsubscribe();
    }

    onCodeSubmit() {
        var code: number = this.codeForm.get('code').value
        this.historyService.updatePostCode(this.currPost.post_id, code).subscribe(resp => {
            this.verified.next(true)
        })
    }

    setPost(post: Post) {
      this.currPost = post
    }

    onDeleteClick(post: Post) {
      if(window.confirm("Are you sure you want to delete the job post: " + post.title + "?"))
      this.historyService.deletePost(post.post_id).subscribe(id => {
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
      this.historyService.setSelectedIdx(i)
    }

    getSelected() {
      return this.historyService.getSelected()
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
      this.historyService.editJobPost(this.currPost.user_id, this.currPost.post_id, this.currPost.title, this.currPost.tags, this.currPost.care_description, this.currPost.start_date, this.currPost.start_time, this.currPost.end_date, this.currPost.end_time, this.currPost.care_type).subscribe(updatedPost => {
        console.log(updatedPost)
      })
    }
}