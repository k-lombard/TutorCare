import {Component, OnInit, ChangeDetectionStrategy} from '@angular/core';
import { Router } from '@angular/router';
import { Observable } from 'rxjs/Observable';
import { User } from '../../models/user.model';
import { select, Store } from '@ngrx/store';
import { AppState } from '../../reducers';
import { getCurrUser, isLoggedIn } from '../../auth/auth.selectors';

@Component({
  selector: 'sidebar',
  templateUrl: './sidebar.component.html',
  styleUrls: ['./sidebar.component.scss']
})
export class SidebarComponent implements OnInit {
  user: User | undefined;
  name = "Account"
  isLoggedIn: boolean = false
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
    this.store
    .pipe(
      select(isLoggedIn)
    ).subscribe(loggedIn => {
      this.isLoggedIn = loggedIn
    })
  }

  public onSidenavClick(): void {
    this.isMenuOpen = false;
  }

  onFindJobsClick() {
    this.router.navigate(['/find-jobs'])
  }

  onMessagesClick() {
    this.router.navigate(['/find-jobs/messages'])
  }

  onMyJobsClick() {
    this.router.navigate(['/find-jobs/my-job-postings'])
  }

  onAppReceivedClick() {
    this.router.navigate(['/find-jobs/applications-received'])
  }

  onActiveJobsClick() {
    this.router.navigate(['/find-jobs/active-jobs'])
  }

  onAppliedClick() {
    this.router.navigate(['/find-jobs/applied-to'])
  }
}
