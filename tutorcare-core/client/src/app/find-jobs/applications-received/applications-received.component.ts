import {Component, OnInit, ChangeDetectionStrategy} from '@angular/core';
import { ActivatedRoute, NavigationEnd, Router } from '@angular/router';
import { select, Store } from '@ngrx/store';
import { DefaultGlobalConfig } from 'ngx-toastr';
import { BehaviorSubject, Subject, Subscription } from 'rxjs';
import { map, switchMap } from 'rxjs/operators';
import { getCurrUser } from 'src/app/auth/auth.selectors';
import { Application } from 'src/app/models/application.model';
import { Post } from 'src/app/models/post.model';
import { User } from 'src/app/models/user.model';
import { AppState } from 'src/app/reducers';
import { GeolocationPositionWithUser } from '../../models/geolocationposition.model';
import { ApplicationsReceivedService } from './applications-received.service';

@Component({
  selector: 'applications-received-component',
  templateUrl: './applications-received.component.html',
  styleUrls: ['./applications-received.component.scss']
})
export class ApplicationsReceivedComponent implements OnInit {

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
    appId!: number
    userType!: string
    private routeSub: Subscription;
    constructor(private router: Router, private appsRec: ApplicationsReceivedService, private store: Store<AppState>, private route: ActivatedRoute) {}


    ngOnInit() {
      this.routeSub = this.route.params.subscribe(params => {
        this.appId = params['id']
        console.log(this.appId)
      });
      if (this.appId) {
        this.appsRec.getApplicationById(this.appId).subscribe(application => {
          console.log(application)
          this.currApp = application
        })
      }
      this.store
        .pipe(
            select(getCurrUser)
        ).subscribe(data =>  {
            this.user = data
            this.userId = this.user.user_id || ""
            this.userType = this.user.user_category
      })
      this.appsRec.getPostsByUserId(this.userId).subscribe(data => {
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

    ngOnDestroy() {
      this.routeSub.unsubscribe();
    }

    onFindCareClick() {
        this.router.navigate(['/find-care'])
    }

    setApp(app: Application) {
      this.currApp = app
    }

    onAcceptApplication(application_id: number, post_id: number, user_id: string, message: string) {
      this.appsRec.acceptApplication(application_id, post_id, user_id, message).subscribe(res => {
        console.log(res)
      })
      this.currApp.accepted = true;
    }




}