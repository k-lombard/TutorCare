import {Component, OnInit, ChangeDetectionStrategy} from '@angular/core';
import { Router } from '@angular/router';
import { Observable } from 'rxjs/Observable';
import { User } from '../models/user.model';
import { select, Store } from '@ngrx/store';
import { AppState } from '../reducers';
import { getCurrUser } from '../auth/auth.selectors';

@Component({
  selector: 'navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.scss'],
  changeDetection: ChangeDetectionStrategy.Default
})
export class NavbarComponent implements OnInit {
  user: User | undefined;
  name = "Account"
  public isMenuOpen: boolean = false;
  constructor(private router: Router, private store: Store<AppState>) {}

  ngOnInit() {
    this.name = "Account"
    this.store
      .pipe(
        select(getCurrUser)
      ).subscribe(data =>  {
        this.user = data
        this.name = (this.user? this.user.first_name + " " + this.user.last_name: "Account")
      })
  }

  public onSidenavClick(): void {
    this.isMenuOpen = false;
  }

  onHomeClick() {
    this.router.navigate(['/home'])
  }

  onFindCareClick() {
    this.router.navigate(['/find-care'])
  }

  onFindJobsClick() {
    this.router.navigate(['/find-jobs'])
  }

  onAboutUsClick() {
    this.router.navigate(['/about-us'])
  }

  onAccountClick() {
    this.router.navigate(['/account'])
  }

  onLoginClick() {
    this.router.navigate(['/login'])
  }

  onSignupClick() {
    this.router.navigate(['/signup'])
  }


}
