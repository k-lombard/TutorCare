import {Component, OnInit, ChangeDetectionStrategy} from '@angular/core';

import { Observable } from 'rxjs';
import { FormControl } from '@angular/forms';
import {MatButtonModule} from '@angular/material/button';
import { LoginComponent } from './auth/login/login.component';
import { Store, select } from '@ngrx/store';
import { AppState } from './reducers';
import { Router } from '@angular/router';
import { isLoggedIn, isLoggedOut } from './auth/auth.selectors';
import { Logout } from './auth/auth.actions';
import { ThemeService } from './theme.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit {
  isLoggedIn$: Observable<boolean> | undefined;

  isLoggedOut$: Observable<boolean> | undefined;
  constructor(private store: Store<AppState>, private router: Router, private theme: ThemeService) {}

  ngOnInit() {
    // this.theme.setTheme()
    this.isLoggedIn$ = this.store
      .pipe(
        select(isLoggedIn)
      );

    this.isLoggedOut$ = this.store
      .pipe(
        select(isLoggedOut)
      );

  }

  logout() {

    this.store.dispatch(new Logout());

  }


}
