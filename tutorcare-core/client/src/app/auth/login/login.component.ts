import {Component, OnInit} from '@angular/core';
import { Observable } from 'rxjs/Observable';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { AuthService } from '../auth.service';
import { tap, map } from 'rxjs/operators';
import { Login } from '../auth.actions';
import { Store } from '@ngrx/store';
import { AppState } from 'src/app/reducers';
import { User } from 'src/app/models/user.model';
import { _getOptionScrollPosition } from '@angular/material/core';
import { ToastrService } from 'ngx-toastr';


@Component({
  selector: 'login-component',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
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
  newpos!: any
  password: string = ""
  emailFC = new FormControl();
  passwordFC = new FormControl();
  accuracy!: number
  latitude!: number
  longitude!: number
  user_id!: string
  loginForm : FormGroup
  constructor(private authService: AuthService, private router:Router, private store: Store<AppState>, private toastr: ToastrService) {}

  validation_messages = {

    'email': [
      { type: 'required', message: 'Email is required' },
      { type: 'email', message: 'Enter a valid email' },
      { type: 'pattern', message: 'Please use a Gatech email' }
    ],
    'password': [
      { type: 'required', message: 'Password is required' },
      { type: 'minlength', message: 'Password must be at least 8 characters long' }
    ]
  }

  ngOnInit() {
    this.createForms();
  }

  onLoginSubmit(value: any) {
    this.email = this.loginForm.get('email').value
    this.password = this.loginForm.get('password').value
    this.loginFunc(this.email, this.password)
  }

  onSignupSubmit() {
    this.router.navigate(['/signup'])
  }

  createForms() {
    this.loginForm = new FormGroup({
      email: new FormControl('', Validators.compose([
        Validators.email,
        Validators.required
      ])),
      password: new FormControl('', Validators.compose([
        Validators.minLength(8),
        Validators.required
      ]))
    })
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
      this.user_id = this.currUser.user_id || ""
      setTimeout(async () => { await this.timeout(10000) }, 100000)
      this.router.navigate(['/home'])
      this.toastr.success("Successfully logged in as " + ((this.currUser.first_name || " ") + " " + (this.currUser.last_name || " ")[0]) + ".", "Success", {closeButton: true, timeOut: 5000, progressBar: true});
    });
    // this.authService.getPosition().subscribe(resp => {
    //   console.log(resp)
    //   this.newpos = resp
    //   this.accuracy = this.newpos.coords.accuracy
    //   this.latitude = this.newpos.coords.latitude
    //   this.longitude = this.newpos.coords.longitude
    //   this.authService.createPosition(this.user_id, this.accuracy, this.latitude, this.longitude).subscribe(resp => {
    //     console.log(resp)
    //   });
    // });
  }
  timeout(ms: number) {
    return new Promise(resolve => setTimeout(resolve, ms));
  }
}
