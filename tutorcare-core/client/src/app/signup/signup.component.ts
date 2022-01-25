import {Component, OnInit, ChangeDetectionStrategy} from '@angular/core';
import {UsersService} from '../users.service';
import {SignupService} from '../signup/signup.service';
import { Observable } from 'rxjs';
import { FormControl } from '@angular/forms';
import {MatButtonModule} from '@angular/material/button';
import { Router } from '@angular/router';

interface Option {
  value: string;
  viewValue: string;
}

@Component({
  selector: 'signup-component',
  templateUrl: './signup.component.html',
  styleUrls: ['./signup.component.scss'],
  changeDetection: ChangeDetectionStrategy.Default
})
export class SignupComponent implements OnInit {
  selectedValue: string | undefined
  userCategory: string = ""
  options: Option[] = [
    {value: 'caregiver-0', viewValue: 'Providing Care'},
    {value: 'careseeker-1', viewValue: 'Seeking Care'},
    {value: 'both-2', viewValue: 'Both'},
  ];
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


  onEmailChange() {
    this.email = this.emailFC.value
  }
  onPasswordChange() {
    this.password = this.passwordFC.value
  }

  onFirstNameChange() {
    this.firstName = this.firstNameFC.value
  }
  onLastNameChange() {
    this.lastName = this.lastNameFC.value
  }

  onSignupSubmit() {
    console.log(this.selectedValue)
    if (this.selectedValue == "caregiver-0") {
      this.userCategory = "caregiver"
    } else if (this.selectedValue == "careseeker-1") {
      this.userCategory = "careseeker"
    } else {
      this.userCategory == "both"
    }
    this.signupFunc(this.firstName, this.lastName, this.email, this.password, this.userCategory)
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

  signupFunc(firstName: string, lastName: string, email: string, password: string, user_category: string) {
    this._signupObservable = this.signupService.signup(firstName, lastName, email, password, user_category);

    this._signupObservable.subscribe((data: any) => {
       this.output = JSON.parse(JSON.stringify(data));
       console.log(data)
    });
    this.router.navigate(['/login'])
  }


}
