import {Component, OnInit, ChangeDetectionStrategy} from '@angular/core';
import {UsersService} from '../users.service';
import {SignupService} from '../signup/signup.service';
import { Observable } from 'rxjs';
import { FormControl } from '@angular/forms';
import {MatButtonModule} from '@angular/material/button';
import { Router } from '@angular/router';

@Component({
  selector: 'signup-component',
  templateUrl: './signup.component.html',
  styleUrls: ['./signup.component.css'],
  changeDetection: ChangeDetectionStrategy.Default
})
export class SignupComponent implements OnInit {
  loading: boolean = false
  _usersObservable: Observable<Object[]> | undefined
  _signupObservable: Observable<Object[]> | undefined
  users: Object | undefined
  output: Object | undefined
  firstName: string = ""
  lastName: string = ""
  email: string = ""
  password: string = ""
  firstNameFC = new FormControl();
  lastNameFC = new FormControl();
  emailFC = new FormControl();
  passwordFC = new FormControl();
  constructor(private router: Router, private usersService: UsersService, private signupService: SignupService) {}

  ngOnInit() {
    this.loading = true
    this.getUsersFunc()
    console.log(this.users)
  }

  onFirstNameChange() {
    this.firstName = this.firstNameFC.value
  } 

  onLastNameChange() {
    this.lastName = this.lastNameFC.value
  } 

  onEmailChange() {
    this.email = this.emailFC.value
  } 
  onPasswordChange() {
    this.password = this.passwordFC.value
  } 

  onSignupSubmit() {
    this.signupFunc(this.firstName, this.lastName, this.email, this.password)
  }

  onLoginSubmit() {
    this.router.navigate(['/login'])
  }

  getUsersFunc() {
    this._usersObservable = this.usersService.getUsers();
 
    this._usersObservable.subscribe((data: any) => {
       this.users = JSON.parse(JSON.stringify(data));
       console.log(data)
    });
  }

  signupFunc(firstName: string, lastName: string, email: string, password: string) {
    this._signupObservable = this.signupService.signup(firstName, lastName, email, password);
 
    this._signupObservable.subscribe((data: any) => {
       this.output = JSON.parse(JSON.stringify(data));
       console.log(data)
    });
    this.router.navigate(['/login'])
  }


}