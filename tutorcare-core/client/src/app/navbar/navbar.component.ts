import {Component, OnInit, ChangeDetectionStrategy} from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css'],
  changeDetection: ChangeDetectionStrategy.Default
})
export class NavbarComponent implements OnInit {
  constructor(private router: Router) {}

  ngOnInit() {
  }

  onHomeClick() {
    this.router.navigate(['/home'])
  }

  onFindCareClick() {
    this.router.navigate(['/find-care'])
  }

  onFindJobsClick() {
    this.router.navigate(['/find-jobs'])
  }

  onAboutUsClick() {
    this.router.navigate(['/about-us'])
  }

  onAccountClick() {
    this.router.navigate(['/account'])
  }

  onLoginClick() {
    this.router.navigate(['/login'])
  }

  onSignupClick() {
    this.router.navigate(['/signup'])
  }


}