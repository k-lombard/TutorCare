import { Component, Inject, OnInit } from "@angular/core";
import { FormBuilder, FormControl, FormGroup } from "@angular/forms";
import { MatDialogRef, MAT_DIALOG_DATA } from "@angular/material/dialog";
import { MatFormFieldControl } from "@angular/material/form-field";
import { select, Store } from "@ngrx/store";
import { Observable } from "rxjs";
import { getCurrUser } from "src/app/auth/auth.selectors";
import { Application } from "src/app/models/application.model";
import { Post } from "src/app/models/post.model";
import { User } from "src/app/models/user.model";
import { AppState } from "src/app/reducers";
import { CreateJobDialog } from "../find-jobs.component";
import { FindJobsService } from "../find-jobs.service";
import { ApplyJobService } from "./apply-job.service";

@Component({
  selector: 'apply-job',
  templateUrl: 'apply-job.html',
  styleUrls: ['./apply-job.component.scss'],
  providers: [
    { provide: MatFormFieldControl, useExisting: ApplyJobDialog }
  ]
})
export class ApplyJobDialog implements OnInit{
  post!: Post
  form!: FormGroup;
  date!: string
  date2!: string
  user!: User
  userId!: string
  message!: string
  postId!: number
  application!: Application
  _applyJobObservable: Observable<Application> | undefined


  constructor(
    public dialogRef: MatDialogRef<CreateJobDialog>,
    @Inject(MAT_DIALOG_DATA) public data: any,
    private fb: FormBuilder, private store: Store<AppState>, private applyJobService: ApplyJobService) {
      this.post = data.post
  }

  ngOnInit() {
    this.form = this.fb.group({
      message: new FormControl(),
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
    this.message = this.form.value.message
    this.postId = this.post.post_id || 0
    this._applyJobObservable = this.applyJobService.applyJob(this.userId, this.postId, this.message)

    this._applyJobObservable.subscribe((data: Application) => {
        console.log(data)
        this.application = data;
    });
    this.dialogRef.close(this.form.value);
  }

  close() {
    this.dialogRef.close();
  }

}
