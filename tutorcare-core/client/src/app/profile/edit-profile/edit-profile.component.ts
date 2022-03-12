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
import { Profile } from 'src/app/models/profile.model';
import { Options, LabelType } from "@angular-slider/ngx-slider";


interface Skill {
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
    loaded = false
    user!: User;
    profile!: Profile
    _editProfileObservable: Observable<User> | undefined
    /*selectedValue!: string
    userCategory: string | undefined
    
    options: Option[] = [
        {value: 'caregiver', viewValue: 'Providing Care'},
        {value: 'careseeker', viewValue: 'Seeking Care'},
        {value: 'both', viewValue: 'Both'},
    ];
    email: string | undefined
    output: string | undefined
    experience: string = ""
    bio: string = ""
    emailFC = new FormControl();
    experienceFC = new FormControl();
    bioFC = new FormControl();
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
    items: Preference[] = [{display: 'Tutoring', value: 'Tutoring'}, {display: 'Baby-sitting', value: 'Baby-sitting'}, {display: 'Dog-sitting', value: 'Dog-sitting'}, {display: 'House-sitting', value: 'House-sitting'}]*/
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
    user_skills: string[] = []

    minValue: number = 20;
    maxValue: number = 30;
    sliderOptions: Options = {
      floor: 10,
      ceil: 100,
      minRange: 10,
      maxRange: 30,
      pushRange: true,
      translate: (value: number, label: LabelType): string => {
        switch (label) {
          case LabelType.Low:
            return "<b>$" + value + "</b>";
          case LabelType.High:
            return "<b>$" + value + "</b>";
          default:
            return "$" + value;
        }
      }
    };
    constructor(private store: Store<AppState>, private router: Router, private route: ActivatedRoute, private editProfileService: EditProfileService) {}

    ngOnInit() {
        this.store
        .pipe(
            select(getCurrUser)
        ).subscribe(data =>  {
            this.user = data
            /*this.emailVal = this.user.email || ""
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
            }*/
        })
        this.editProfileService.getProfile(this.user.user_id).subscribe( data => {
          this.profile = data
          this.user_skills = this.skillsToArray(this.profile.skills)
          this.parseRateRange(this.profile.rate_range)
          console.log(this.user_skills)
      })
      this.loaded = true
      this.user_skills.values()
    }

    skillsToArray(skill_list: string) {
      var skill_array: string[] = []
      skill_list.split(',').forEach(parsedString => {
        var skill = this.all_skills.find(item => item.value == parsedString)
        if (skill) {
          skill_array.push(skill.value)
        }
      });
      return skill_array
    }

    skillsToString(skill_list: string[]) {
      var skill_string = ""
      skill_list.forEach((item) => {
        skill_string += item + ","
      })
      return skill_string
    }

    parseRateRange(rate: string) {
      var str = rate.split(",")
      this.minValue = (Number.isNaN(parseInt(str[0])) ? this.minValue : parseInt(str[0]))
      this.maxValue = (Number.isNaN(parseInt(str[1])) ? this.maxValue : parseInt(str[1]))
    }

    /*onEmailChange() {
        this.email = this.emailFC.value
    }

    onExperienceChange() {
        this.experience = this.emailFC.value
    }

    onBioChange() {
        this.bio = this.emailFC.value
    }*/

    onCancelSubmit() {
        this.router.navigate(['account'])
    }

    onSave() {
        /*if (this.selectedValue == 'caregiver-1') {
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
        }*/
        console.log("saved")
        console.log(this.profile)
        this.profile.rate_range = "$" + this.minValue.toString() + " - $" + this.maxValue.toString();
        this.profile.skills = this.skillsToString(this.user_skills) || ""
        this.editProfileFunc(this.user.user_id, this.profile)
    }

    checkCovid() {
      console.log(this.profile)
      this.profile.covid19 = !this.profile.covid19
    }
    checkSmoker() {
      this.profile.smoker = !this.profile.smoker
    }
    checkCpr() {
      this.profile.cpr = !this.profile.cpr
    }
    checkFirstAid() {
      this.profile.first_aid = !this.profile.first_aid
    }

    editProfileFunc(user_id: string | undefined, newProfile: Profile) {
      console.log(newProfile)
        this._editProfileObservable = this.editProfileService.editProfile(user_id, newProfile);
        this._editProfileObservable.subscribe((data: Profile) => {
            console.log(data)
            //this.profile = data
            //this.user = data;
            //this.user.refresh_token = this.refresh_token
            //this.user.access_token = this.access_token
            //this.store.dispatch(new Login({user: data}));
        });
        this.router.navigate([`/profile/${this.user.user_id}`]).then(() => {
          window.location.reload();
        });
    }

}
