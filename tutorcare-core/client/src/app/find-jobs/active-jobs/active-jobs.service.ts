import { Injectable } from '@angular/core';
import {Http} from '@angular/http';
import {environment} from '../../../environments/environment';
import { Observable } from 'rxjs/Observable';
import { map } from 'rxjs/operators';
import { GeolocationPositionWithUser } from '../../models/geolocationposition.model';
import { Post } from '../../models/post.model';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Application } from 'src/app/models/application.model';


@Injectable()
export class ActiveJobsService {
  results:Object[];
  _posts: Post[] | undefined;
  _applications: Application[] | undefined;
  _output: any[] | undefined;
  headers = new HttpHeaders({
    'Content-Type': 'application/json'
  });
  constructor(private http: HttpClient) {
    this.results = []
  }

  getActiveJobsByUserId(user_id: string): Observable<Post[]> {
    let url = `/api/posts/active-jobs/${user_id}`;
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
