import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/Observable';
import { catchError, map } from 'rxjs/operators';
import { Post } from '../../models/post.model';
import { HttpClient, HttpErrorResponse, HttpHeaders } from '@angular/common/http';
import { Application } from 'src/app/models/application.model';
import { throwError } from 'rxjs';
import { ToastrService } from 'ngx-toastr';
import { PostCode } from '../../models/postcode.model';


@Injectable()
export class HistoryService {
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

getVerificationCode(post_id: number) {
    let url = `/api/post_codes/${post_id}`;
    return new Observable((observer: any) => {
      this.http.get(url)
         .pipe(map((res: any) => res))
         .subscribe((data: PostCode) => {
            observer.next(data);
            observer.complete();
         });
  });
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

updateJobPostCaregiver(post_id: number, post: Post): Observable<Post> {
    let url = `/api/posts/${post_id}`;
    return new Observable((observer: any) => {
        this.http.put<any>(url, JSON.stringify({
            "caregiver_completed": true,
            "post_id": post_id,
            "care_type": post.care_type,
            "user_id": post.user_id,
            "title": post.title,
            "care_description": post.care_description
        }), {headers: this.headers})
            .pipe(map((res: any) => res))
            .subscribe((data: any) => {
   
               observer.next(data);
               observer.complete();
            });
     });
}

updateJobPostPoster(post_id: number, post: Post): Observable<Post> {
    let url = `/api/posts/${post_id}`;
    return new Observable((observer: any) => {
        this.http.put<any>(url, JSON.stringify({
            "poster_completed": true,
            "post_id": post_id,
            "care_type": post.care_type,
            "user_id": post.user_id,
            "title": post.title,
            "care_description": post.care_description
        }), {headers: this.headers})
            .pipe(map((res: any) => res))
            .subscribe((data: any) => {
   
               observer.next(data);
               observer.complete();
            });
     });
}

updatePostCode(post_id: number, code: number): Observable<PostCode> {
    let url = `/api/post_codes/${post_id}`
    return new Observable((observer: any) => {
        this.http.put<any>(url, JSON.stringify({
            "verified": true,
            "post_id": post_id
        }), {headers: this.headers})
            .pipe(map((res: any) => res))
            .subscribe((data: any) => {
   
               observer.next(data);
               observer.complete();
            });
     });
}

async getPostByIdPromise(post_id: number): Promise<PostCode> {
  let url = `/api/posts/${post_id}`;
  return await this.http.get<PostCode>(url, {headers: this.headers}).toPromise()
}

async getPostsByUserIdCompleted(user_id: string): Promise<any> {
    let url = `/api/posts/user/completed/${user_id}`;
    return await this.http.get<PostCode>(url, {headers: this.headers}).toPromise()
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
