import {Component, OnInit, ChangeDetectionStrategy} from '@angular/core';
import { Store, select } from '@ngrx/store';
import { AppState } from '../reducers';
import { User } from '../models/user.model';
import { getCurrUser } from '../auth/auth.selectors';
import { Router, ActivatedRoute, NavigationExtras} from '@angular/router';

@Component({
    selector: 'account-component',
    templateUrl: './account.component.html',
    styleUrls: ['./account.component.scss']
})
export class AccountComponent implements OnInit {
    user!: User;
    first_name: string | undefined
    last_name: string |undefined
    rate: number = 4.5
    user_type: string | undefined
    constructor(private store: Store<AppState>, private router: Router, private route: ActivatedRoute) {}

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
    }

    onEditProfileClick() {
        this.router.navigate(['edit-profile'], {relativeTo: this.route})
    }



}