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

@Component({
    selector: 'account-component',
    templateUrl: './account.component.html',
    styleUrls: ['./account.component.scss']
})
export class AccountComponent implements OnInit {
    url = this.router.url
    user!: User;
    first_name: string | undefined
    last_name: string |undefined
    rate: number = 4.5
    user_type: string | undefined
    exp!: string
    bio!: string
    careseeker: boolean = false
    caregiver: boolean = false
    both: boolean = false
    constructor(private store: Store<AppState>, private router: Router, private route: ActivatedRoute, private authService: AuthService, private toastr: ToastrService) {}

    ngOnInit() {
        this.store
        .pipe(
            select(getCurrUser)
        ).subscribe(data =>  {
            this.user = data
            this.first_name = (this.user ? this.user.first_name : "")
            this.last_name = (this.user ? this.user.last_name : "")
            this.user_type = this.user.user_category? this.user.user_category.charAt(0).toUpperCase() + this.user.user_category.substring(1) : ""
            this.exp = this.user.experience || ""
            this.bio = this.user.bio || ""
            
        })
        if (this.user_type == "Caregiver") {
          this.caregiver = true
        }
        else if (this.user_type == "Careseeker") {
          this.careseeker = true
        } else {
          this.both = true
        }
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
