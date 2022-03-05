import {Component, OnInit, ChangeDetectionStrategy} from '@angular/core';
import { Store, select } from '@ngrx/store';
import { AppState } from '../reducers';
import { User } from '../models/user.model';
import { getCurrUser } from '../auth/auth.selectors';
import { Router, ActivatedRoute, NavigationExtras} from '@angular/router';
import { AuthService } from '../auth/auth.service';
import { ToastrService } from 'ngx-toastr';
import { tap } from 'rxjs/operators';
import { Logout } from '../auth/auth.actions';
import { Profile } from '../models/profile.model';
import { ProfileService } from './profile.service';
import { Subscription } from 'rxjs';

interface Badges {
  value: string
  description: string
  matIconString: string
}

@Component({
    selector: 'profile-component',
    templateUrl: './profile.component.html',
    styleUrls: ['./profile.component.scss']
})
export class ProfileComponent implements OnInit {
    url = this.router.url
    user!: User;
    currentUser!: User;
    profile!: Profile;
    /*first_name: string | undefined
    last_name: string |undefined
    user_type: string | undefined*/
    all_badges: Badges[] = [
      {value: 'verified', description: 'Verified University Email', matIconString:'verified_user'},
      {value: '20jobs', description: '20 Completed Jobs', matIconString:'whatshot'}
    ];
    user_badges: Badges[] = []
    private routeSub: Subscription
    constructor(private store: Store<AppState>, private router: Router, private route: ActivatedRoute, private authService: AuthService, private toastr: ToastrService, private profileService: ProfileService) {}

    ngOnInit() {
      this.routeSub = this.route.params.subscribe(params => {
        this.profileService.getUserByUserId(params['id']).subscribe((data: User) => {
          this.user = data
        })
        this.profileService.getProfileByUserId(params['id']).subscribe((data: Profile) => {
          this.profile = data
          this.setBadges(this.profile.badge_list)
        })
      })
      this.store
      .pipe(
          select(getCurrUser)
      ).subscribe(data =>  {
          this.currentUser = data
          /*this.first_name = (this.user ? this.user.first_name : "")
          this.last_name = (this.user ? this.user.last_name : "")
          this.user_type = this.user.user_category? this.user.user_category.charAt(0).toUpperCase() + this.user.user_category.substring(1) : ""*/
      })
    }

    setBadges(badge_list: string) {
      badge_list.split(',').forEach(parsedString => {
        var badge = this.all_badges.find(item => item.value == parsedString)
        if (badge) {
          this.user_badges.push(badge)
        }
      });
    }

    onEditProfileClick() {
      this.router.navigate(['profile/edit-profile'])//, {relativeTo: this.route.parent})
    }

    logoutFunc() {
      this.authService.logout()
      .pipe(
        tap(user => {
          this.store.dispatch(new Logout());
        })
      )
      .subscribe(resp => {
        console.log(resp)
        this.router.navigate(['/home'])
        this.toastr.success("Successfully logged out.", "Success", {closeButton: true, timeOut: 5000, progressBar: true});
      });
    }
}