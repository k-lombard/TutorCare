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

interface Skill {
  display: string,
  value: string
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
      {value: 'verified', description: 'Verified GA Tech Email', matIconString:'verified_user'},
      {value: '20jobs', description: '20 Completed Jobs', matIconString:'whatshot'}
    ];
    user_badges: Badges[] = []
    all_skills: Skill[] = [
      {display: 'Tutoring', value: 'Tutoring'},
      {display: 'Baby-sitting', value: 'Baby-sitting'},
      {display: 'Dog-sitting', value: 'Dog-sitting'},
      {display: 'House-sitting', value: 'House-sitting'},
      {display: 'Math', value: 'Math'},
      {display: 'Chemistry', value: 'Chemistry'},
      {display: 'Biology', value: 'Biology'},
      {display: 'Calculus', value: 'Calculus'},
      {display: 'Physics', value: 'Physics'},
      {display: 'Algebra', value: 'Algebra'},
      {display: 'Geometry', value: 'Geometry'},
      {display: 'Computer Science', value: 'Computer_Science'},
      {display: 'Mechanical Engineering', value: 'Mechanical_Engineering'},
      {display: 'Neuroscience', value: 'Neuroscience'},
      {display: 'Chemical Engineering', value: 'Chemical_Engineering'},
      {display: 'Industrial Engineering', value: 'Industrial_Engineering'},
      {display: 'Aeronautical Engineering', value: 'Aeuronautical_Engineering'},
      {display: 'Industrial Engineering', value: 'Industrial_Engineering'},
      {display: 'Business', value: 'Business'},
      {display: 'Linear Algebra', value: 'Linear_Algebra'},
      {display: 'Multivariable Calculus', value: 'Multivariable_Calculus'},
      {display: 'Ages 0-2', value: 'Ages_0-2'},
      {display: 'Ages 3-6', value: 'Ages_3-6'},
      {display: 'Ages 7-10', value: 'Ages_7-10'},
      {display: 'Ages 11-14', value: 'Ages_11-14'},
      {display: 'Ages 15-17', value: 'Ages_15-17'},
  ]
  user_skills: Skill[] = []
    private routeSub: Subscription
    constructor(private store: Store<AppState>, private router: Router, private route: ActivatedRoute, private authService: AuthService, private toastr: ToastrService, private profileService: ProfileService) {}

    ngOnInit() {
      this.routeSub = this.route.params.subscribe(params => {
        console.log(params['id'])
        this.profileService.getUserByUserId(params['id']).subscribe((data: User) => {
          this.user = data
        })
        this.profileService.getProfileByUserId(params['id']).subscribe((data: Profile) => {
          this.profile = data
          //this.profile.badge_list = "verified"
          this.setBadges(this.profile.badge_list)
          this.user_skills = this.skillsToArray(this.profile.skills)
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

    skillsToArray(skill_list: string) {
      var skill_array: Skill[] = []
      skill_list.split(',').forEach(parsedString => {
        var skill = this.all_skills.find(item => item.value == parsedString)
        if (skill) {
          skill_array.push(skill)
        }
      });
      return skill_array
    }

    onEditProfileClick() {
      this.router.navigate([`/profile/${this.currentUser.user_id}/edit-profile`])//, {relativeTo: this.route})
    }

    onLogoutClick() {
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

    onSettingsClick() {
      this.router.navigate(['settings'])
    }
}