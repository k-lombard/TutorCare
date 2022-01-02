import {Component, OnInit, ChangeDetectionStrategy} from '@angular/core';
import {UsersService} from './users.service';
import {LoginService} from './login.service';
import { Observable } from 'rxjs';
import { FormControl } from '@angular/forms';
import {MatButtonModule} from '@angular/material/button';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
  changeDetection: ChangeDetectionStrategy.Default
})
export class AppComponent implements OnInit {
  loading: boolean = false
  _usersObservable: Observable<Object[]> | undefined
  _loginObservable: Observable<Object[]> | undefined
  users: Object | undefined
  output: Object | undefined
  email: string = ""
  password: string = ""
  emailFC = new FormControl();
  passwordFC = new FormControl();
  constructor(private usersService: UsersService, private loginService: LoginService) {}

  ngOnInit() {
    this.loading = true
    this.getUsersFunc()
    console.log(this.users)
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

  ngOnChanges() {

  }

  getUsersFunc() {
    this._usersObservable = this.usersService.getUsers();
 
    this._usersObservable.subscribe((data: any) => {
       this.users = JSON.parse(JSON.stringify(data));
       console.log(data)
    });
  }

  loginFunc(email: string, password: string) {
    this._loginObservable = this.loginService.login(email, password);
 
    this._loginObservable.subscribe((data: any) => {
       this.output = JSON.parse(JSON.stringify(data));
       console.log(data)
    });
  }


}