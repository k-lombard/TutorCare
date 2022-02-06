import {Component, OnInit, ChangeDetectionStrategy} from '@angular/core';
import { Store, select } from '@ngrx/store';
import { AppState } from '../../reducers';
import { User } from '../../models/user.model';
import { getCurrUser } from '../../auth/auth.selectors';
import { Router, ActivatedRoute, NavigationExtras} from '@angular/router';
import { FormControl, FormGroup } from '@angular/forms';
import { EditProfileService } from './edit-profile.service';
import { Observable, pipe } from 'rxjs';
import { tap } from 'rxjs/operators';
import { Login, Logout } from 'src/app/auth/auth.actions';
import { AuthService } from 'src/app/auth/auth.service';
import { ToastrService } from 'ngx-toastr';

interface Option {
    value: string;
    viewValue: string;
}

interface Preference {
  display: string,
  value: string
}

@Component({
    selector: 'edit-profile-component',
    templateUrl: './edit-profile.component.html',
    styleUrls: ['./edit-profile.component.scss']
})
export class EditProfileComponent implements OnInit {
    public form: FormGroup = new FormGroup({})
    selectedValue!: string
    userCategory: string | undefined
    _editProfileObservable: Observable<User> | undefined
    options: Option[] = [
        {value: 'caregiver-0', viewValue: 'Providing Care'},
        {value: 'careseeker-1', viewValue: 'Seeking Care'},
        {value: 'both-2', viewValue: 'Both'},
    ];
    email: string | undefined
    output: string | undefined
    experience: string = ""
    bio: string = ""
    emailFC = new FormControl();
    experienceFC = new FormControl();
    bioFC = new FormControl();
    user!: User;
    first_name: string | undefined
    last_name: string |undefined
    rate: number = 4.5
    user_type: string | undefined
    user_id: string | undefined
    expVal!: string
    bioVal!: string
    emailVal!: string
    catVal!: string
    access_token!: string
    refresh_token!: string
    selectedPreferences: string[] = []
    preferenceString: string = ""
    items: Preference[] = [{display: 'Tutoring', value: 'Tutoring'}, {display: 'Baby-sitting', value: 'Baby-sitting'}, {display: 'Dog-sitting', value: 'Dog-sitting'}, {display: 'House-sitting', value: 'House-sitting'}]
    constructor(private store: Store<AppState>, private router: Router, private route: ActivatedRoute, private editProfileService: EditProfileService) {}

    ngOnInit() {
        this.store
        .pipe(
            select(getCurrUser)
        ).subscribe(data =>  {
            this.user = data
            this.emailVal = this.user.email || ""
            this.catVal = this.user.user_category || ""
            this.expVal = this.user.experience || ""
            this.bioVal = this.user.bio || ""
            this.access_token = this.user.access_token || ""
            this.refresh_token = this.user.refresh_token || ""
            if (this.catVal == 'caregiver') {
                this.selectedValue = 'caregiver-0'
            } else if (this.catVal == 'careseeker') {
                this.selectedValue = 'careseeker-1'
            } else if (this.catVal == 'both') {
                this.selectedValue = 'both-2'
            }
            if (this.user.preferences) {
              var prefCopy: string[] = []
              for (let str of this.user.preferences.split(" ")) {
                var tempStr = (str.slice(0,1).toUpperCase() + str.slice(1))
                prefCopy.push(tempStr)
              }
              this.selectedPreferences = prefCopy
            }
        })
    }

    onEmailChange() {
        this.email = this.emailFC.value
    }

    onExperienceChange() {
        this.experience = this.emailFC.value
    }

    onBioChange() {
        this.bio = this.emailFC.value
    }

    onCancelSubmit() {
        this.router.navigate(['account'])
    }

    onSave() {
        if (this.selectedValue == 'caregiver-0') {
          this.userCategory = "caregiver"
        } else if (this.selectedValue == 'careseeker-1') {
          this.userCategory = "careseeker"
        } else if (this.selectedValue == 'both-2') {
          this.userCategory = "both"
        }
        for (let val of this.selectedPreferences) {
          if (this.preferenceString === "") {
            this.preferenceString = val.toLowerCase()
          } else {
            this.preferenceString = this.preferenceString + " " + val.toLowerCase()
          }
        }
        this.editProfileFunc(this.user.user_id, this.emailVal, this.expVal, this.userCategory, this.bioVal, this.user.password, this.preferenceString)
    }

    editProfileFunc(user_id: string | undefined, email: string | undefined, experience: string | undefined, user_category: string | undefined, bio: string | undefined, password: string | undefined, preferences: string | undefined) {
        this._editProfileObservable = this.editProfileService.editProfile(user_id, email, experience, user_category, bio, password, preferences);

        this._editProfileObservable.subscribe((data: User) => {
            this.user = data;
            this.user.refresh_token = this.refresh_token
            this.user.access_token = this.access_token
            this.store.dispatch(new Login({user: data}));
        });
        this.router.navigate(['/account'])
    }

}
