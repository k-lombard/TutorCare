import {Component, OnInit, ChangeDetectionStrategy, Inject} from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { GeolocationPositionWithUser } from '../models/geolocationposition.model';
import { Post } from '../models/post.model';
import { FindJobsService } from './find-jobs.service';
import {MatDialog, MatDialogConfig, MatDialogRef, MAT_DIALOG_DATA} from '@angular/material/dialog';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';
import { MatFormFieldControl } from '@angular/material/form-field';
import { select, Store } from '@ngrx/store';
import { AppState } from '../reducers';
import { getCurrUser, isLoggedIn } from '../auth/auth.selectors';
import { User } from '../models/user.model';
import { ThisReceiver } from '@angular/compiler';
import { Observable } from 'rxjs';
import { ApplyJobDialog } from './apply-job/apply-job.component';
import { ToastrService } from 'ngx-toastr';


interface FilterOption {
    value: string;
    viewValue: string;
  }
@Component({
  selector: 'find-jobs-component',
  templateUrl: './find-jobs.component.html',
  styleUrls: ['./find-jobs.component.scss']
})
export class FindJobsComponent implements OnInit {
    selectedValue: string | undefined
    rate: number = 4.5
    userCategory: string = ""
    posts: Post[] = []
    div1: boolean = false;
    search: string =""
    userType!: string
    user!: User
    userId!: string
    isLoggedIn!: boolean
    filter_options: FilterOption[] = [
        {value: 'tutoring-0', viewValue: 'Type: Tutoring'},
        {value: 'babysitting-1', viewValue: 'Type: Babysitting'},
        {value: 'other-2', viewValue: 'Type: Other'}
    ];
    constructor(private router: Router, private findJobs: FindJobsService, public dialog: MatDialog, private route: ActivatedRoute, private store: Store<AppState>, private toastr: ToastrService) {}

    openDialog() {
      if (this.isLoggedIn) {
        const dialogConfig = new MatDialogConfig();

        dialogConfig.disableClose = false;
        // dialogConfig.autoFocus = true;
        dialogConfig.data = {
          id: 1,
          title: 'Create Job Posting',
        };
        dialogConfig.height = "90%"
        dialogConfig.width = "90%"
        const dialogRef = this.dialog.open(CreateJobDialog, dialogConfig);

        dialogRef.afterClosed().subscribe(
          data => console.log("Dialog output:", data)
        );
      } else {
        this.toastr.error("You must be logged in to do this.", "Error", {closeButton: true, timeOut: 5000, progressBar: true})
      }
    }

    openApplyDialog(post: Post) {
      if (this.isLoggedIn) {
        const dialogConfig = new MatDialogConfig();

        dialogConfig.disableClose = false;
        // dialogConfig.autoFocus = true;
        dialogConfig.data = {
          id: 1 as number,
          title: 'Apply to this Job Posting' as string,
          post: post as Post,
          posts: this.posts as Post[]
        };
        dialogConfig.height = "90%"
        dialogConfig.width = "90%"
        const dialogRef = this.dialog.open(ApplyJobDialog, dialogConfig);

        dialogRef.afterClosed().subscribe(
          data => {
            console.log("Dialog output:", data)
          }
        );
        console.log(dialogConfig.data.posts)
        this.posts = dialogConfig.data.posts
      } else {
        this.toastr.error("You must be logged in to do this.", "Error", {closeButton: true, timeOut: 5000, progressBar: true})
      }
    }

    div1Function(){
      this.div1=!this.div1;
    }

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
            this.userId = this.user.user_id || ""
            this.userType = this.user.user_category
      })
        this.findJobs.getPosts().subscribe(data => {
            this.posts = data
            var postsCopy: Post[] = []
            console.log(this.posts)
            for (var post of this.posts) {
              post.tagList = post.tags.split(" ")
              postsCopy.push(post)
            }
            this.posts = postsCopy
        })

    }
    onFindCareClick() {
        this.router.navigate(['/find-care'])
    }


}
@Component({
  selector: 'create-job',
  templateUrl: 'create-job.html',
  styleUrls: ['./create-job.component.scss'],
  providers: [
    { provide: MatFormFieldControl, useExisting: CreateJobDialog }
  ]
})
export class CreateJobDialog implements OnInit{
  post!: Post
  form!: FormGroup;
  date!: string
  date2!: string
  type_care!: string
  job_desc!: string
  date_of_job!: string
  picker!: string
  endTimeFC = new FormControl()
  startTimeFC = new FormControl()
  start_time!: string
  end_time!: string
  user!: User
  userId!: string
  month!: number
  dayStr!: string
  monthStr!: string
  day!: number
  title!: string
  tags!: string

  _createJobObservable: Observable<Post> | undefined
  minDate = new Date()
  maxDate = new Date().setMonth(new Date().getMonth() + 3)
  filter_options: FilterOption[] = [
    {value: 'tutoring-0', viewValue: 'Type: Tutoring'},
    {value: 'babysitting-1', viewValue: 'Type: Babysitting'},
    {value: 'other-2', viewValue: 'Type: Other'}
];


  constructor(
    public dialogRef: MatDialogRef<CreateJobDialog>,
    @Inject(MAT_DIALOG_DATA) public data: any,
    private fb: FormBuilder, private store: Store<AppState>, private findJobs: FindJobsService) {

  }

  onStartTimeChange() {
    this.start_time = this.startTimeFC.value
  }
  onEndTimeChange() {
    this.end_time = this.endTimeFC.value
  }

  ngOnInit() {
    this.form = this.fb.group({
      type_care: new FormControl(),
      job_title: new FormControl(),
      job_desc: new FormControl(),
      job_tags: new FormControl(),
      picker: new FormControl(new Date()),
      start_time: new FormControl(),
      end_time: new FormControl(),
      posts: new FormControl()
    });
    this.store
        .pipe(
            select(getCurrUser)
        ).subscribe(data =>  {
            this.user = data
            this.userId = this.user.user_id || ""
    })

  }

  save() {
    console.log(this.form.value)
    this.job_desc = this.form.value.job_desc
    if (this.form.value.type_care == "tutoring-0") {
      this.type_care = "tutoring"
    } else if (this.form.value.type_care == "babysitting-1") {
      this.type_care = "baby-sitting"
    } else if (this.form.value.type_care == "other-2") {
      this.type_care = "other"
    }
    this.month = this.form.value.start_time.getMonth() + 1
    this.day = this.form.value.start_time.getDate()
    console.log(this.day)
    if (this.month < 10) {
      this.monthStr = '0' + this.month
    } else {
      this.monthStr = this.month.toString()
    }
    if (this.day < 10) {
      this.dayStr = '0' + this.day
    } else {
      this.dayStr = this.day.toString()
    }
    this.date_of_job = this.form.value.start_time.getFullYear() + '-' + this.monthStr + '-' + this.dayStr
    this.start_time = this.form.value.start_time.getHours() + ':' + this.form.value.start_time.getMinutes()
    this.end_time = this.form.value.end_time.getHours() + ':' + this.form.value.end_time.getMinutes()
    this.title = this.form.value.job_title
    this.tags = this.form.value.job_tags
    console.log(this.start_time)
    console.log(this.form.value.start_time.getMonth())
    console.log(this.form.value.start_time.getDay())
    this._createJobObservable = this.findJobs.createPost(this.userId, this.title, this.job_desc, this.tags, this.type_care, this.date_of_job, this.start_time, this.end_time)

    this._createJobObservable.subscribe((data2: Post) => {
        this.post = data2
        if (this.data.posts && this.data.posts instanceof Array) {
          this.form.value.posts = this.data.posts.concat(this.post)
        }

    });
    this.dialogRef.close(this.form.value);
  }

  close() {
    this.dialogRef.close();
  }

}

