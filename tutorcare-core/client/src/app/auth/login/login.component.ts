import {Component, OnInit, ChangeDetectionStrategy} from '@angular/core';
import { Observable } from 'rxjs/Observable';
import { FormControl } from '@angular/forms';
import {MatButtonModule} from '@angular/material/button';
import { Router } from '@angular/router';
import { AuthService } from '../auth.service';
import { tap, map } from 'rxjs/operators';
import { noop } from 'rxjs';
import { Login } from '../auth.actions';
import { Store } from '@ngrx/store';
import { AppState } from 'src/app/reducers';
import Swal from 'sweetalert2'
import { User } from 'src/app/models/user.model';
import { _getOptionScrollPosition } from '@angular/material/core';
import { ToastrService } from 'ngx-toastr';


@Component({
  selector: 'login-component',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  loading: boolean = false
  _usersObservable: Observable<Object[]> | undefined
  _loginObservable: Observable<Object[]> | undefined
  users: Object | undefined
  output: Object | undefined
  currUser!: User
  email: string = ""
  pos: string | undefined
  password: string = ""
  emailFC = new FormControl();
  passwordFC = new FormControl();
  constructor(private authService: AuthService, private router:Router, private store: Store<AppState>, private toastr: ToastrService) {}

  ngOnInit() {
  }

  onEmailChange() {
    this.email = this.emailFC.value
  } 
  onPasswordChange() {
    this.password = this.passwordFC.value
  } 

  onLoginSubmit() {
    this.loginFunc(this.email, this.password)
  }

  onSignupSubmit() {
    this.router.navigate(['/signup'])
  }

  // getUsersFunc() {
  //   this._usersObservable = this.usersService.getUsers();
 
  //   this._usersObservable.subscribe((data: any) => {
  //      this.users = JSON.parse(JSON.stringify(data));
  //      console.log(data)
  //   });
  // }

  loginFunc(email: string, password: string) {
    this.authService.login(this.email, this.password)
    .pipe(
      tap(user => {
        this.store.dispatch(new Login({user}));
      })
    )
    .subscribe(resp => {
      console.log(resp)
      this.currUser = resp
      setTimeout(async () => { await this.timeout(10000) }, 100000)
      this.router.navigate(['/home'])
      this.toastr.success("Successfully logged in as " + ((this.currUser.first_name || " ") + " " + (this.currUser.last_name || " ")[0]) + ".", "Success", {closeButton: true, timeOut: 5000, progressBar: true});
    });
    this.authService.getPosition().subscribe(resp => {
      console.log(resp)
    })
  }
  timeout(ms: number) {
    return new Promise(resolve => setTimeout(resolve, ms));
  }
}