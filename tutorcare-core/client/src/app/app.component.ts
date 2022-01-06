import {Component, OnInit, ChangeDetectionStrategy} from '@angular/core';
import {UsersService} from './users.service';
import {LoginService} from './login/login.service';
import { Observable } from 'rxjs';
import { FormControl } from '@angular/forms';
import {MatButtonModule} from '@angular/material/button';
import { LoginComponent } from './login/login.component';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
  changeDetection: ChangeDetectionStrategy.Default
})
export class AppComponent implements OnInit {
  constructor(private usersService: UsersService, private loginService: LoginService) {}

  ngOnInit() {
    
  }
}
