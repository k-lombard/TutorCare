import { Injectable } from '@angular/core';
import {Http} from '@angular/http';
import {environment} from '../../../environments/environment';
import { Observable } from 'rxjs/Observable';
import { catchError, map } from 'rxjs/operators';
import { GeolocationPositionWithUser } from '../../models/geolocationposition.model';
import { Post } from '../../models/post.model';
import { HttpClient, HttpErrorResponse, HttpHeaders } from '@angular/common/http';
import { Application } from 'src/app/models/application.model';
import { throwError } from 'rxjs';
import { ToastrService } from 'ngx-toastr';


@Injectable()
export class MyJobPostingsService {
  results:Object[];
  _posts: Post[] | undefined;
  _applications: Application[] | undefined;
  _output: any[] | undefined;
  _post: Post | undefined
  selectedIdx!: number
  _post_id: string | undefined
  headers = new HttpHeaders({
    'Content-Type': 'application/json'
  });
  constructor(private http: HttpClient, private toastr: ToastrService) {
    this.results = []
  }

  getPostsByUserId(user_id: string): Observable<Post[]> {
    let url = `/api/posts/user/${user_id}`;
    return new Observable((observer: any) => {
       this.http.get(url)
           .pipe(map((res: any) => res.posts))
           .subscribe((data: Post[]) => {
              this._posts = data

              observer.next(this._posts);
              observer.complete();
           });
    });
 }

  getApplicationsByPostId(post_id: number): Observable<Application[]> {
    let url = `/api/applications/post/${post_id}`;
    return new Observable((observer: any) => {
      this.http.get(url)
         .pipe(map((res: any) => res.applications))
         .subscribe((data: Application[]) => {
            this._applications= data

            observer.next(this._applications);
            observer.complete();
         });
  });
}

deletePost(post_id: number): Observable<string> {
  let url = `/api/posts/${post_id}`;
    return new Observable((observer: any) => {
      this.http.delete<any>(url)
         .pipe(map((res: any) => res))
         .subscribe((data: string) => {
            this._post_id = data

            observer.next(this._post_id);
            observer.complete();
         });
  });
}

editJobPost(user_id: string, post_id: number, title: string, tags: string, care_description: string, start_date: string, start_time: string, end_date:string, end_time: string, care_type: string): Observable<Post> {
  let url = `/api/posts/${post_id}`;
    return new Observable((observer: any) => {
      this.http.put<any>(url, JSON.stringify({
        "user_id": user_id,
        "title": title,
        "tags": tags,
        "care_description": care_description,
        "start_date": start_date,
        "start_time": start_time,
        "end_date": end_date,
        "end_time": end_time,
        "care_type": care_type
    }), {headers: this.headers})
         .pipe(map((res: any) => res),
         catchError((err: HttpErrorResponse) => {
          this.toastr.error("Invalid date or time format: " + err.status, "Error", {closeButton: true, timeOut: 5000, progressBar: true});
          return throwError(err)
        })
         )
         .subscribe((data: Post) => {
            observer.next(data);
            observer.complete();
         },
         error => {return throwError(error)}
       );
  });
}

setSelectedIdx(i: number) {
  this.selectedIdx = i
}

getSelected() {
  return this.selectedIdx
}

getPostById(post_id: number): Observable<Post> {
  let url = `/api/posts/${post_id}`;
    return new Observable((observer: any) => {
      this.http.get(url)
         .pipe(map((res: any) => res))
         .subscribe((data: Post) => {
            this._post = data

            observer.next(this._post);
            observer.complete();
         });
  });
}

acceptApplication(application_id?: number, post_id?: number, user_id?: string, message?: string): Observable<Post> {
  // let url = `${environment.serverUrl}/api/signup/`;
  let url = `/api/applications/accept/${application_id}`;
  return new Observable((observer: any) => {
     this.http.put<any>(url, JSON.stringify({
         "application_id": application_id,
         "post_id": post_id,
         "user_id": user_id,
         "message": message,
         "accepted": true
     }), {headers: this.headers})
         .pipe(map((res: any) => res))
         .subscribe((data: any) => {
            this._output = data

            observer.next(this._output);
            observer.complete();
         });
  });
}


}
