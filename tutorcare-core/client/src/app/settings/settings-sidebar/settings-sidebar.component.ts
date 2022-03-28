import {Component, OnInit, ChangeDetectionStrategy, Input} from '@angular/core';
import { Router } from '@angular/router';
import { Observable } from 'rxjs/Observable';
import { User } from '../../models/user.model';
import { select, Store } from '@ngrx/store';
import { AppState } from '../../reducers';
import { getCurrUser, isLoggedIn } from '../../auth/auth.selectors';

@Component({
  selector: 'settings-sidebar',
  templateUrl: './settings-sidebar.component.html',
  styleUrls: ['./settings-sidebar.component.scss']
})
export class SettingsSidebarComponent implements OnInit {
  @Input() currPage: string;
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

  onMyAccountClick() {
    this.router.navigate(['/settings'])
  }

  onPastJobsClick() {
    console.log("CLICK")
    this.router.navigate(['/settings/past-jobs'])
  }
  
}
