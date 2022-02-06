import {Component, OnInit, ChangeDetectionStrategy} from '@angular/core';
import {UsersService} from '../../users.service';
import { Observable} from 'rxjs';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { ParentErrorStateMatcher, PasswordValidator } from '../validators/password.validator';
import { VerifyService } from './verify.service';
interface Option {
  value: string;
  viewValue: string;
}

@Component({
  selector: 'verify-component',
  templateUrl: './verify.component.html',
  styleUrls: ['./verify.component.scss'],
  changeDetection: ChangeDetectionStrategy.Default
})

export class VerifyComponent implements OnInit {
  _usersObservable: Observable<Object[]> | undefined
  _signupObservable: Observable<Object[]> | undefined
  _verifyObservable: Observable<string> | undefined
  users: Object | undefined
  output: Object | undefined
  accountDetailsForm: FormGroup;
  matchingPasswordsGroup: FormGroup;
  emailCodeForm: FormGroup;
  resendCodeForm: FormGroup;
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
    private verifyService: VerifyService,
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
    'emailVerificationCode': [
      { type: 'required', message: 'Verification code is required' },
      { type: 'minlength', message: 'Email code must be at least 5 characters long' },
    ],
    'email2': [
      { type: 'required', message: 'Email is required' },
      { type: 'email', message: 'Enter a valid email' },
      { type: 'pattern', message: 'Please use a Gatech email' },
      { type: 'exists', message: 'That email already has an account'}
    ]
  }



  ngOnInit() {
    this.createForms();
  }

  createForms() {
    this.resendCodeForm = new FormGroup({
      email: new FormControl('', Validators.compose([
        Validators.email,
        // Validators.pattern('^.*gatech.edu.*$'),
        Validators.required
      ]))
    })

    // Email Form validations
    this.emailCodeForm = this.fb.group({
      emailVerificationCode: new FormControl('', Validators.compose([
        Validators.minLength(5),
        Validators.required
      ])),
      email2: new FormControl('', Validators.compose([
        Validators.email,
        Validators.required
      ]))
    })
  }

  onVerifySubmit(value: any) {
    this._verifyObservable = this.verifyService.verifyCode(this.emailCodeForm.get('email2').value, parseInt(this.emailCodeForm.get('emailVerificationCode').value))
    this._verifyObservable.subscribe(
      (data: any) => {
        this.output = data;
        console.log(data)
    })
    this.router.navigate(['/login'])
  }

  onResendEmailSubmit() {
    console.log("Resend Email")
    this.verifyService.resendCode(this.resendCodeForm.get('email').value).subscribe(data => {
      console.log(data)
    })
  }
}
