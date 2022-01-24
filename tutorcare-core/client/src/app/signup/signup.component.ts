import {Component, OnInit, ChangeDetectionStrategy} from '@angular/core';
import {UsersService} from '../users.service';
import {SignupService} from '../signup/signup.service';
import { Observable } from 'rxjs';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { ParentErrorStateMatcher, PasswordValidator } from './validators/password.validator';

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
  loading: boolean = false
  _usersObservable: Observable<Object[]> | undefined
  _signupObservable: Observable<Object[]> | undefined
  users: Object | undefined
  output: Object | undefined
  accountDetailsForm: FormGroup;
  matchingPasswordsGroup: FormGroup;
  constructor(
    private router: Router, 
    private usersService: UsersService, 
    private signupService: SignupService,
    private fb: FormBuilder) {}
  parentErrorStateMatcher = new ParentErrorStateMatcher();

  account_validation_messages = {
    'name': [
      { type: 'required', message: 'Name is required' },
      { type: 'maxlength', message: 'Name cannot be more than 25 characters long' },
      { type: 'pattern', message: 'Your name must contain only numbers and letters' }
    ],
    'email': [
      { type: 'required', message: 'Email is required' },
      { type: 'email', message: 'Enter a valid email' },
      { type: 'pattern', message: 'Please use a Gatech email' }
    ],
    'confirm_password': [
      { type: 'required', message: 'Confirm password is required' },
      { type: 'areEqual', message: 'Password mismatch' }
    ],
    'password': [
      { type: 'required', message: 'Password is required' },
      { type: 'minlength', message: 'Password must be at least 5 characters long' },
      { type: 'pattern', message: 'Your password must contain at least one uppercase, one lowercase, and one number' }
    ]
  }

  

  ngOnInit() {
    this.createForms();
  }

  createForms() {
    // matching passwords validation
    this.matchingPasswordsGroup = new FormGroup({
      password: new FormControl('', Validators.compose([
        Validators.minLength(5),
        Validators.required,
        Validators.pattern('^(?=.*?[A-Z])(?=.*?[a-z])(?=.*?[0-9])(?=.*?[#?!@$%^&*-]).+$')/*^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])[a-zA-Z0-9]+$*/
      ])),
      confirm_password: new FormControl('', Validators.required)
    }, (formGroup: FormGroup) => {
      return PasswordValidator.areEqual(formGroup);
    });

    // account form validations
    this.accountDetailsForm = this.fb.group({
      firstName: new FormControl('', Validators.compose([
       Validators.maxLength(25),
       Validators.pattern('^[a-zA-Z0-9_]*$'),
       Validators.required
      ])),
      lastName: new FormControl('', Validators.compose([
        Validators.maxLength(25),
        Validators.pattern('^[a-zA-Z0-9_]*$'),
        Validators.required
       ])),
      email: new FormControl('', Validators.compose([
        Validators.email,
        //Validators.pattern('^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+.[a-zA-Z0-9-.]+$'), //Regex Email Pattern
        Validators.pattern('^.*gatech.edu.*$'),
        Validators.required
      ])),
      matchingPasswords: this.matchingPasswordsGroup,
    })
  }

  onSignupSubmit(value: any){
    console.log(value);
    this.signupFunc(
      this.accountDetailsForm.get('firstName').value, 
      this.accountDetailsForm.get('lastName').value, 
      this.accountDetailsForm.get('email').value, 
      this.accountDetailsForm.get('matchingPasswords').get('password').value,
      "caregiver" //user_catagory temp
      )
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