import {Component, OnInit, ChangeDetectionStrategy} from '@angular/core';
import { ActivatedRoute, NavigationEnd, Router } from '@angular/router';
import { select, Store } from '@ngrx/store';
import { DefaultGlobalConfig, ToastrService } from 'ngx-toastr';
import { BehaviorSubject, Subject, Subscription } from 'rxjs';
import { map, switchMap } from 'rxjs/operators';
import { getCurrUser } from 'src/app/auth/auth.selectors';
import { Application } from 'src/app/models/application.model';
import { Post } from 'src/app/models/post.model';
import { User } from 'src/app/models/user.model';
import { AppState } from 'src/app/reducers';
import { GeolocationPositionWithUser } from '../../models/geolocationposition.model';
import { MyJobPostingsService } from './my-job-postings.service';

@Component({
  selector: 'my-job-postings-component',
  templateUrl: './my-job-postings.component.html',
  styleUrls: ['./my-job-postings.component.scss']
})
export class MyJobPostingsComponent implements OnInit {

    public acceptSubject: Subject<boolean> = new BehaviorSubject<boolean>(false);
    public acceptActive = this.acceptSubject.asObservable();
    selectedValue: string | undefined
    userCategory: string = ""
    user!: User
    userId!: string
    posts!: Post[]
    currPost!: Post
    menuVisible: boolean
    locs!: GeolocationPositionWithUser[]
    mySubscription!: any
    postId!: number
    userType!: string
    editable: boolean = true
    private routeSub: Subscription;
    constructor(private router: Router, private myJobs: MyJobPostingsService, private store: Store<AppState>, private route: ActivatedRoute, private toastr: ToastrService) {}


    ngOnInit() {
      this.routeSub = this.route.params.subscribe(params => {
        this.postId = parseInt(params['id'])
        console.log(this.postId)
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
        this.myJobs.getPostById(this.postId).subscribe(post => {
          console.log(post)
          this.currPost = post
        })
      }
      this.myJobs.getPostsByUserId(this.userId).subscribe(data => {
        this.posts = data
        if (this.posts) {
          for (let i = 0; i < this.posts?.length; i++) {
            if (this.posts[i].post_id === this.postId) {
              this.myJobs.setSelectedIdx(i)
            }
          }
        }
        var postsCopy: Post[] = []
        if (this.posts) {
          for (var post of this.posts) {
            post.tagList = post.tags.split(" ")
            postsCopy.push(post)
          }
          this.posts = postsCopy
          console.log(this.posts)
       }
      })
    }

    ngOnDestroy() {
      this.routeSub.unsubscribe();
    }

    onFindCareClick() {
        this.router.navigate(['/find-care'])
    }

    setPost(post: Post) {
      this.currPost = post
    }

    onDeleteClick(post_id: number) {
      this.myJobs.deletePost(post_id).subscribe(id => {
        console.log("deleted post with id: " + id)
        this.currPost = undefined
        this.toastr.success("Post successfully deleted with id: "+ id, "Success", {closeButton: true, timeOut: 5000, progressBar: true});
        var postsCopy: Post[] = []
        for (let pos of this.posts) {
          if (pos.post_id != post_id) {
            postsCopy.push(pos)
          }
        }
        this.posts = postsCopy
      })
    }

    onEditClick() {
      this.editable = false
      console.log(this.editable)
    }

    setSelected(i: number) {
      this.myJobs.setSelectedIdx(i)
    }

    getSelected() {
      return this.myJobs.getSelected()
    }

    back() {
      this.currPost= undefined
    }

    backToMenu() {
      this.currPost = undefined
      this.menuVisible = true
    }

    onSaveClick() {
      this.editable = true
      this.myJobs.editJobPost(this.currPost.user_id, this.currPost.post_id, this.currPost.title, this.currPost.tags, this.currPost.care_description, this.currPost.date_of_job, this.currPost.start_time, this.currPost.end_time, this.currPost.care_type).subscribe(updatedPost => {
        console.log(updatedPost)
      })
    }




}
