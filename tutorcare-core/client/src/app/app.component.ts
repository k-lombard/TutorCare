import {Component, OnInit} from '@angular/core';
import {LoginService} from './login.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {

  title: any;

  constructor(private hw: LoginService) {}

  ngOnInit() {
    this.hw.getTitle()
      .subscribe((data: any) => this.title = data.title);

    console.log(this.title);
  }

}