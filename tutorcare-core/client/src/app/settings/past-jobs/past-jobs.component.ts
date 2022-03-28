import {Component, OnInit} from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';
import { Store, select } from '@ngrx/store';
import { ToastrService } from 'ngx-toastr';
import { getCurrUser, isLoggedIn } from '../../auth/auth.selectors';
import { AuthService } from '../../auth/auth.service';
import { Profile } from '../../models/profile.model';
import { User } from '../../models/user.model';
import { ProfileService } from '../../profile/profile.service';
import { AppState } from '../../reducers';
import { PastJobsService } from './past-jobs.service';


@Component({
  selector: 'past-jobs-component',
  templateUrl: './past-jobs.component.html',
  styleUrls: ['./past-jobs.component.scss']
})
export class PastJobsComponent implements OnInit {
  
  constructor(private store: Store<AppState>, private router: Router, private route: ActivatedRoute, private authService: AuthService, private toastr: ToastrService, private settingsService:PastJobsService ) {}


  userProfile!: Profile
  user!: User
  menuVisible: boolean
  mainCol: boolean
  isLoggedIn!: boolean
  ngOnInit() {
    this.store
      .pipe(
        select(isLoggedIn)
      ).subscribe(data2 => {
        this.isLoggedIn = data2
      })
      this.store
        .pipe(
            select(getCurrUser)
        ).subscribe(data =>  {
          this.user = data
          this.settingsService.getProfileByUserId(this.user.user_id).subscribe((data: Profile) => {
            this.userProfile = data
          })
        })
  }

  search(){
    console.log("Searching")
  }

  backToMenu() {
    this.mainCol = false
    this.menuVisible = true
  }
}
