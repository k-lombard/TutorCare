import { Injectable } from '@angular/core';
import {Http} from '@angular/http';
import {environment} from '../../environments/environment';
import { Observable } from 'rxjs/Observable';
import { map } from 'rxjs/operators';
import { GeolocationPositionWithUser } from '../models/geolocationposition.model';
import { Post } from '../models/post.model';
import { HttpClient, HttpHeaders } from '@angular/common/http';


@Injectable()
export class FindJobsService {
  results:Object[];
  _posts: Post[] | undefined;
  _output: any[] | undefined;
  headers = new HttpHeaders({
    'Content-Type': 'application/json'
  });
  constructor(private http: HttpClient) {
    this.results = []
  }

  getPosts(): Observable<Post[]> {
    let url = `/api/posts/active`;
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

 createPost(userId: string, jobTitle: string, careDesc: string, tags: string, careType: string, start_date: string, start_time: string, end_date: string, end_time: string, repeat: string): Observable<Post> {
  let url = `/api/posts/`;
  return new Observable((observer: any) => {
    // @ts-ignore
     this.http.post<any>(url, JSON.stringify({
         "user_id": userId,
         "title": jobTitle,
         "care_description": careDesc,
         "tags": tags,
         "care_type": careType,
         "start_date": start_date,
         "start_time": start_time,
         "end_date": end_date,
         "end_time": end_time,
         "job_repeat": repeat,

     }), {headers: this.headers})
         .pipe(map((res: any) => res))
         .subscribe((data: any) => {
            this._output = data

            observer.next(this._output);
            observer.complete();

         });
  });
}

//  getPosition(): Observable<any> {
//   return new Observable(observer => {
//     window.navigator.geolocation.getCurrentPosition(position => {
//       observer.next(position);
//       observer.complete();
//     },
//       error => observer.error(error));
//   });
// }

}
