import {Component, OnInit, ChangeDetectionStrategy, Inject} from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { MatDialog, MatDialogConfig, MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { MatFormFieldControl } from '@angular/material/form-field';
import { ActivatedRoute, NavigationEnd, Router } from '@angular/router';
import { select, Store } from '@ngrx/store';
import { DefaultGlobalConfig, ToastrService } from 'ngx-toastr';
import { BehaviorSubject, Observable, Subject, Subscription } from 'rxjs';
import { map, switchMap } from 'rxjs/operators';
import { getCurrUser, isLoggedIn } from 'src/app/auth/auth.selectors';
import { Application } from 'src/app/models/application.model';
import { Post } from 'src/app/models/post.model';
import { User } from 'src/app/models/user.model';
import { AppState } from 'src/app/reducers';
import { GeolocationPositionWithUser } from '../../models/geolocationposition.model';
import { FindJobsService } from '../find-jobs.service';
import { MyJobPostingsService } from './my-job-postings.service';
import { NgModule, CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';
import { DateValidator } from '../date.validator';

@Component({
  selector: 'my-job-postings-component',
  templateUrl: './my-job-postings.component.html',
  styleUrls: ['./my-job-postings.component.scss']
})
export class MyJobPostingsComponent implements OnInit {

    public acceptSubject: Subject<boolean> = new BehaviorSubject<boolean>(false);
    public acceptActive = this.acceptSubject.asObservable();
    selectedValue: string | undefined
    userCategory: string = ""
    user!: User
    userId!: string
    posts!: Post[]
    currPost!: Post
    menuVisible: boolean
    mainCol: boolean
    locs!: GeolocationPositionWithUser[]
    mySubscription!: any
    postId!: number
    userType!: string
    editable: boolean = true
    private routeSub: Subscription;
    //constructor(private router: Router, private myJobs: MyJobPostingsService, private store: Store<AppState>, private route: ActivatedRoute, private toastr: ToastrService) {}

    constructor(private router: Router, public dialog: MatDialog, private route: ActivatedRoute, private store: Store<AppState>, private toastr: ToastrService, private myJobs: MyJobPostingsService) {}

    openDialog(post: Post) {
        const dialogConfig = new MatDialogConfig();
        console.log(post)
        dialogConfig.disableClose = false;
        // dialogConfig.autoFocus = true;
        dialogConfig.data = {
          id: post.post_id,
          title: 'Edit Job Posting',
          post: post
        };
        dialogConfig.height = "90%"
        dialogConfig.width = "80%"
        const dialogRef = this.dialog.open(EditJobDialog, dialogConfig);

        dialogRef.afterClosed().subscribe(
          data => console.log("Dialog output:", data)
        );
    }

    ngOnInit() {
      this.routeSub = this.route.params.subscribe(params => {
        this.postId = parseInt(params['id'])
        console.log(this.postId)
      });
      this.store
        .pipe(
            select(getCurrUser)
        ).subscribe(data =>  {
            this.user = data
            this.userId = this.user.user_id || ""
            this.userType = this.user.user_category
      })
      if (this.postId && (!this.posts || this.posts.length === 0)) {
        this.myJobs.getPostById(this.postId).subscribe(post => {
          console.log(post)
          this.currPost = post
        })
      }
      this.myJobs.getPostsByUserId(this.userId).subscribe(data => {
        this.posts = data
        if (this.posts) {
          for (let i = 0; i < this.posts?.length; i++) {
            if (this.posts[i].post_id === this.postId) {
              this.myJobs.setSelectedIdx(i)
            }
          }
        }
        var postsCopy: Post[] = []
        if (this.posts) {
          for (var post of this.posts) {
            post.tagList = post.tags.split(" ")
            postsCopy.push(post)
          }
          this.posts = postsCopy
          console.log(this.posts)
       }
      })
    }

    ngOnDestroy() {
      this.routeSub.unsubscribe();
    }

    onFindCareClick() {
        this.router.navigate(['/find-care'])
    }

    setPost(post: Post) {
      this.currPost = post
    }

    onDeleteClick(post: Post) {
      if(window.confirm("Are you sure you want to delete the job post: " + post.title + "?"))
      this.myJobs.deletePost(post.post_id).subscribe(id => {
        console.log("deleted post with id: " + id)
        this.currPost = undefined
        this.toastr.success("Post successfully deleted with id: "+ id, "Success", {closeButton: true, timeOut: 5000, progressBar: true});
        var postsCopy: Post[] = []
        for (let pos of this.posts) {
          if (pos.post_id != post.post_id) {
            postsCopy.push(pos)
          }
        }
        this.posts = postsCopy
      })
    }

    onEditClick() {
      this.editable = false
      console.log(this.editable)
    }

    setSelected(i: number) {
      this.myJobs.setSelectedIdx(i)
    }

    getSelected() {
      return this.myJobs.getSelected()
    }

    back() {
      this.currPost= undefined
    }

    backToMenu() {
      this.mainCol = false
      this.menuVisible = true
    }

    onSaveClick() {
      this.editable = true
      this.myJobs.editJobPost(this.currPost.user_id, this.currPost.post_id, this.currPost.title, this.currPost.tags, this.currPost.care_description, this.currPost.start_date, this.currPost.start_time, this.currPost.end_date, this.currPost.end_time, this.currPost.care_type).subscribe(updatedPost => {
        console.log(updatedPost)
      })
    }
}

interface FilterOption {
  value: string;
  viewValue: string;
}

interface Tag {
  display: string,
  value: string
}

@Component({
  selector: 'edit-job',
  templateUrl: 'edit-job.html',
  styleUrls: ['./edit-job.component.scss'],
  providers: [
    { provide: MatFormFieldControl, useExisting: EditJobDialog }
  ]
})
export class EditJobDialog implements OnInit{
  post!: Post
  form!: FormGroup
  dateGroup!: FormGroup
  date!: string
  date2!: string
  type_care!: string
  job_desc!: string
  end_date!: string
  start_date!: string
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
  selectedTags: string[] = []
  tagString: string = ""
  items: Tag[] =
  [
    {display: 'Tutoring', value: 'Tutoring'},
    {display: 'Baby-sitting', value: 'Baby-sitting'},
    {display: 'Dog-sitting', value: 'Dog-sitting'},
    {display: 'House-sitting', value: 'House-sitting'},
    {display: 'Math', value: 'Math'},
    {display: 'Chemistry', value: 'Chemistry'},
    {display: 'Biology', value: 'Biology'},
    {display: 'Calculus', value: 'Calculus'},
    {display: 'Physics', value: 'Physics'},
    {display: 'Algebra', value: 'Algebra'},
    {display: 'Geometry', value: 'Geometry'},
    {display: 'Computer Science', value: 'Computer_Science'},
    {display: 'Mechanical Engineering', value: 'Mechanical_Engineering'},
    {display: 'Neuroscience', value: 'Neuroscience'},
    {display: 'Chemical Engineering', value: 'Chemical_Engineering'},
    {display: 'Industrial Engineering', value: 'Industrial_Engineering'},
    {display: 'Aeronautical Engineering', value: 'Aeuronautical_Engineering'},
    {display: 'Industrial Engineering', value: 'Industrial_Engineering'},
    {display: 'Business', value: 'Business'},
    {display: 'Linear Algebra', value: 'Linear_Algebra'},
    {display: 'Multivariable Calculus', value: 'Multivariable_Calculus'},
    {display: 'Ages 0-2', value: 'Ages_0-2'},
    {display: 'Ages 3-6', value: 'Ages_3-6'},
    {display: 'Ages 7-10', value: 'Ages_7-10'},
    {display: 'Ages 11-14', value: 'Ages_11-14'},
    {display: 'Ages 15-17', value: 'Ages_15-17'},
  ]

  validation_messages = {
    'care_type': [
      { type: 'required', message: 'Type is required' }
    ],
    'job_title': [
      { type: 'required', message: 'Title is required' },
      { type: 'maxlength', message: 'Title cannot be more than 50 characters long' }
    ],
    'job_desc': [
      { type: 'required', message: 'Job decription is required' },
      { type: 'maxlength', message: 'Description cannot be more than 1023 characters long' }
    ],
    'start_time': [
      { type: 'required', message: 'Starting date and time are required' },
    ],
    'end_time': [
      { type: 'required', message: 'Ending date and time are required' },
      { type: 'dateLessThan', message: 'Ending time must be after starting time'} //BUG: does not show
    ]
    
  }

  _editJobObservable: Observable<Post> | undefined
  minDate1 = new Date()
  minDate2 = new Date()
  maxDate = new Date()
  filter_options: FilterOption[] = [
    {value: 'tutoring-0', viewValue: 'Type: Tutoring'},
    {value: 'baby-sitting-1', viewValue: 'Type: Babysitting'},
    {value: 'other-2', viewValue: 'Type: Other'}
];


  constructor(
    public dialogRef: MatDialogRef<EditJobDialog>,
    @Inject(MAT_DIALOG_DATA) public data,
    private fb: FormBuilder, private store: Store<AppState>, private myJobs: MyJobPostingsService) {
      this.post = this.data.post
      this.minDate1.setHours((new Date()).getHours() + 1)
      this.minDate2.setHours((new Date()).getHours() + 2)
      this.maxDate.setMonth((new Date()).getMonth() + 3)
    }

  ngOnInit() {
    this.dateGroup = new FormGroup({
      start_time: new FormControl('', Validators.compose([
        Validators.required
      ])),
      end_time: new FormControl('', Validators.compose([
        Validators.required
      ]))
    }, (formGroup: FormGroup) => {
        return DateValidator.dateLessThan(formGroup);
    })

    this.form = this.fb.group({
      type_care: new FormControl('', Validators.compose([
        Validators.required
      ])),
      job_title: new FormControl('', Validators.compose([
        Validators.required,
        Validators.maxLength(50)
      ])),
      job_desc: new FormControl('', Validators.compose([
        Validators.required,
        Validators.maxLength(1023)
      ])),
      dateGroup: this.dateGroup
    });
    this.store
        .pipe(
            select(getCurrUser)
        ).subscribe(data =>  {
            this.user = data
            this.userId = this.user.user_id || ""
    })
    if (this.post.care_type == "tutoring") {
      this.form.get('type_care').setValue("tutoring-0")
    } else if (this.post.care_type == "babysitting") {
      this.form.get('type_care').setValue("baby-sitting-1")
    } else if (this.post.care_type == "other") {
      this.form.get('type_care').setValue("other-2")
    }
    this.form.get('job_title').setValue(this.post.title)
    this.form.get('job_desc').setValue(this.post.care_description)
    this.selectedTags = this.post.tagList
    //this.form.get('start_time').setValue(this.post.start_time)
    //this.form.get('end_time').setValue(this.post.end_time)
  }

  onStartTimeChange() {
    this.start_time = this.startTimeFC.value
  }

  onEndTimeChange() {
    this.end_time = this.endTimeFC.value
  }

  save() {
    var start_time = new Date(this.form.get('start_time').value)
    var end_time = new Date(this.form.get('end_time').value)
    console.log(this.form.value)
    this.job_desc = this.form.value.job_desc
    if (this.form.value.type_care == "tutoring-0") {
      this.type_care = "tutoring"
    } else if (this.form.value.type_care == "babysitting-1") {
      this.type_care = "baby-sitting"
    } else if (this.form.value.type_care == "other-2") {
      this.type_care = "other"
    }
    var month = start_time.getMonth() + 1
    var day = start_time.getDate()
    console.log(this.day)
    var monthStr
    if (month < 10) {
      monthStr = '0' + month
    } else {
      monthStr = month.toString()
    }
    var dayStr
    if (day < 10) {
      dayStr = '0' + this.day
    } else {
      dayStr = day.toString()
    }
    this.start_date = start_time.getFullYear() + '-' + monthStr + '-' + dayStr

    var month2 = end_time.getMonth() + 1
    var day2 = end_time.getDate()
    console.log(this.day)
    var monthStr2
    if (month2 < 10) {
      monthStr2 = '0' + month2
    } else {
      monthStr2 = month2.toString()
    }
    var dayStr2
    if (day2 < 10) {
      dayStr2 = '0' + day2
    } else {
      dayStr2 = day2.toString()
    }
    this.end_date = end_time.getFullYear() + '-' + monthStr2 + '-' + dayStr2
    
    for (let val of this.selectedTags) {
      if (this.tagString === "") {
        this.tagString = val.toLowerCase()
      } else {
        this.tagString = this.tagString + " " + val.toLowerCase()
      }
    }
    this.tags = this.tagString

    this.start_time = start_time.getHours() + ':' + start_time.getMinutes()
    this.end_time = end_time.getHours() + ':' + end_time.getMinutes()
    this.title = this.form.value.job_title
    
    console.log(this.userId, this.post.post_id, this.title, this.tags, this.job_desc, this.start_date, this.start_time, this.end_date, this.end_time, this.type_care)

    this._editJobObservable = this.myJobs.editJobPost(this.userId, this.post.post_id, this.title, this.tags, this.job_desc, this.start_date, this.start_time, this.end_date, this.end_time, this.type_care)

    this._editJobObservable.subscribe((data2: Post) => {
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
