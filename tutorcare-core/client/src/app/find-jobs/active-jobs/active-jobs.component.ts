import {Component, OnInit, ChangeDetectionStrategy} from '@angular/core';
import { NavigationEnd, Router } from '@angular/router';
import { select, Store } from '@ngrx/store';
import { DefaultGlobalConfig } from 'ngx-toastr';
import { BehaviorSubject, Subject } from 'rxjs';
import { map, switchMap } from 'rxjs/operators';
import { getCurrUser } from 'src/app/auth/auth.selectors';
import { Application } from 'src/app/models/application.model';
import { Post } from 'src/app/models/post.model';
import { User } from 'src/app/models/user.model';
import { AppState } from 'src/app/reducers';
import { GeolocationPositionWithUser } from '../../models/geolocationposition.model';
import { ActiveJobsService } from './active-jobs.service';

@Component({
  selector: 'active-jobs-component',
  templateUrl: './active-jobs.component.html',
  styleUrls: ['./active-jobs.component.scss']
})
export class ActiveJobsComponent implements OnInit {

    public acceptSubject: Subject<boolean> = new BehaviorSubject<boolean>(false);
    public acceptActive = this.acceptSubject.asObservable();
    selectedValue: string | undefined
    userCategory: string = ""
    user!: User
    userId!: string
    posts!: Post[]
    currApp!: Application
    locs!: GeolocationPositionWithUser[]
    mySubscription!: any
    constructor(private router: Router, private activeJobs: ActiveJobsService, private store: Store<AppState>) {}


    ngOnInit() {
      this.store
        .pipe(
            select(getCurrUser)
        ).subscribe(data =>  {
            this.user = data
            this.userId = this.user.user_id || ""
      })
      this.activeJobs.getActiveJobsByUserId(this.userId).subscribe(data => {
        this.posts = data
        var postsCopy: Post[] = []
        for (var post of this.posts) {
          post.tagList = post.tags.split(" ")
          postsCopy.push(post)
        }
        this.posts = postsCopy
        console.log(this.posts)
      })


    }
}
