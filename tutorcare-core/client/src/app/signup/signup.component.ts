import {Component, OnInit, ChangeDetectionStrategy} from '@angular/core';
import {UsersService} from '../users.service';
import {SignupService} from '../signup/signup.service';
import { Observable} from 'rxjs';
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
  _usersObservable: Observable<Object[]> | undefined
  _signupObservable: Observable<Object[]> | undefined
  _verifyObservable: Observable<string> | undefined
  users: Object | undefined
  output: Object | undefined
  accountDetailsForm: FormGroup;
  matchingPasswordsGroup: FormGroup;
  emailCodeForm: FormGroup;
  userCategory!: string
  selectedValue!: string
  options: Option[] = [
    {value: 'caregiver', viewValue: 'Provide Care'},
    {value: 'careseeker', viewValue: 'Find Care'},
    {value: 'both', viewValue: 'Both'},
  ];
  hidden = true
  constructor(
    private router: Router,
    private usersService: UsersService,
    private signupService: SignupService,
    private fb: FormBuilder) {}
  parentErrorStateMatcher = new ParentErrorStateMatcher();

  validation_messages = {
    'name': [
      { type: 'required', message: 'Name is required' },
      { type: 'maxlength', message: 'Name cannot be more than 25 characters long' },
      { type: 'pattern', message: 'Your name must contain only numbers and letters' }
    ],
    'email': [
      { type: 'required', message: 'Email is required' },
      { type: 'email', message: 'Enter a valid email' },
      { type: 'pattern', message: 'Please use a Gatech email' },
      { type: 'exists', message: 'That email already has an account'}
    ],
    'confirm_password': [
      { type: 'required', message: 'Confirm password is required' },
      { type: 'areEqual', message: 'Password mismatch' }
    ],
    'password': [
      { type: 'required', message: 'Password is required' },
      { type: 'minlength', message: 'Password must be at least 8 characters long' },
      { type: 'pattern', message: 'Your password must contain at least one uppercase, one lowercase, and one number' }
    ],
    'emailCode': [
      { type: 'required', message: 'Password is required' },
      { type: 'minlength', message: 'Email code must be at least 5 characters long' },
    ],
    'city': [
      { type: 'required', message: 'City is required' },
      { type: 'email', message: 'Enter a valid city' },
      { type: 'pattern', message: 'Please enter a valid city' },
    ],
    'zipcode': [
      { type: 'required', message: 'Zipcode is required' },
      { type: 'email', message: 'Enter a valid zipcode' },
      { type: 'pattern', message: 'Please enter a valid zipcode' },
    ],
    'address': [
      { type: 'required', message: 'Address is required' },
      { type: 'email', message: 'Enter a valid address' },
      { type: 'pattern', message: 'Please enter a valid address' },
    ]
  }



  ngOnInit() {
    this.createForms();
  }

  createForms() {
    // Matching passwords validation
    this.matchingPasswordsGroup = new FormGroup({
      password: new FormControl('', Validators.compose([
        Validators.minLength(8),
        Validators.required,
        Validators.pattern('^(?=.*?[A-Z])(?=.*?[a-z])(?=.*?[0-9])(?=.*?[#?!@$%^&*-]).+$')/*^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])[a-zA-Z0-9]+$*/
      ])),
      confirm_password: new FormControl('', Validators.required)
    }, (formGroup: FormGroup) => {
      return PasswordValidator.areEqual(formGroup);
    });

    // Account form validations
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
        // Validators.pattern('^.*gatech.edu.*$'),
        Validators.required
      ])),
      city: new FormControl('', Validators.compose([
        Validators.maxLength(25),
        Validators.pattern('^[a-zA-Z_ ]*$'),
        Validators.required
       ])),
      zipcode: new FormControl('', Validators.compose([
        Validators.maxLength(20),
        Validators.pattern('^[0-9_]*$'),
        Validators.required
       ])),
      address: new FormControl('', Validators.compose([
        Validators.maxLength(40),
        Validators.pattern('^[a-zA-Z0-9_ ]*$'),
        Validators.required
       ])),
      matchingPasswords: this.matchingPasswordsGroup
    })

    // Email Form validations
    this.emailCodeForm = this.fb.group({
      emailVerificationCode: new FormControl('', Validators.compose([
        Validators.minLength(5),
        Validators.required
      ]))
    })
  }


  onLoginSubmit() {
    this.router.navigate(['/login'])
  }

  onSignupSubmit(value: any){
    if (this.selectedValue == null) {
      this.selectedValue = "caregiver"
    }
    this.signupFunc(
      this.accountDetailsForm.get('firstName').value,
      this.accountDetailsForm.get('lastName').value,
      this.accountDetailsForm.get('email').value,
      this.accountDetailsForm.get('matchingPasswords').get('password').value,
      this.selectedValue,
      this.accountDetailsForm.get('city').value,
      this.accountDetailsForm.get('zipcode').value,
      this.accountDetailsForm.get('address').value,
      )
  }

  getUsersFunc() {
    this._usersObservable = this.usersService.getUsers();
    this._usersObservable.subscribe((data: any) => {
       this.users = JSON.parse(JSON.stringify(data));
       console.log(data)
    });
  }

  signupFunc(firstName: string, lastName: string, email: string, password: string, user_category: string, city: string, zipcode: string, address: string) {
    this._signupObservable = this.signupService.signup(firstName, lastName, email, password, user_category, city, zipcode, address);
    console.log(this._signupObservable)

    this._signupObservable.subscribe(
      (data: any) => {
        this.output = JSON.parse(JSON.stringify(data));
        console.log(data)
        this.hidden = false
      },
      err => {
        console.log("There was an error", err)
        this.hidden = true;
      }
    );
    this.router.navigate(['/verify'])
  }
}
