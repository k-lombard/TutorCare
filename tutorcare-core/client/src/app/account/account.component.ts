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
import { ProfileComponent } from '../profile/profile.component';
import { AccountService } from './account.service';

interface Badges {
  value: string
  description: string
  matIconString: string
}

@Component({
    selector: 'account-component',
    templateUrl: './account.component.html',
    styleUrls: ['./account.component.scss']
})
export class AccountComponent implements OnInit {
    url = this.router.url
    user!: User;
    profile!: Profile;
    first_name: string | undefined
    last_name: string |undefined
    user_type: string | undefined
    exp!: string
    bio!: string
    all_badges: Badges[] = [
      {value: 'verified', description: 'Verified University Email', matIconString:'verified_user'},
      {value: '20jobs', description: '20 Completed Jobs', matIconString:'whatshot'}
    ];
    user_badges: Badges[] = []
    constructor(private store: Store<AppState>, private router: Router, private route: ActivatedRoute, private authService: AuthService, private toastr: ToastrService, private accountService: AccountService) {}

    ngOnInit() {
        this.store
        .pipe(
            select(getCurrUser)
        ).subscribe(data =>  {
            this.user = data
            this.first_name = (this.user ? this.user.first_name : "")
            this.last_name = (this.user ? this.user.last_name : "")
            this.user_type = this.user.user_category? this.user.user_category.charAt(0).toUpperCase() + this.user.user_category.substring(1) : ""
        })
        this.accountService.getProfile(this.user.user_id).subscribe( data => {
            this.profile = data
            /*if (this.user.status) {
              this.profile.badge_list += "verified" + ","
            }
            this.profile.jobs_completed = 20
            if (this.profile.jobs_completed >= 20) {
              this.profile.badge_list += "20jobs" + ","
            }*/
            // Just for testing badge. Badge strings need to be added when they occur. Ex: when verified email, "verified" + "," gets appended then not here
            this.setBadges(this.profile.badge_list)
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
        this.router.navigate(['edit-profile'], {relativeTo: this.route})
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





