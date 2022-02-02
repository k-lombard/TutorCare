import { Injectable } from '@angular/core';
import {Http} from '@angular/http';
import {environment} from '../../../environments/environment';
import { Observable } from 'rxjs/Observable';
import { map } from 'rxjs/operators';
import { GeolocationPositionWithUser } from '../../models/geolocationposition.model';
import { Post } from '../../models/post.model';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Application } from 'src/app/models/application.model';
import { Chatroom } from 'src/app/models/chatroom.model';


@Injectable()
export class ApplicationsReceivedService {
  results:Object[];
  selectedIdxMap: Map<number, number> = new Map<number,number>()
  _posts: Post[] | undefined;
  _applications: Application[] | undefined;
  _output: any[] | undefined;
  _application: Application | undefined
  headers = new HttpHeaders({
    'Content-Type': 'application/json'
  });
  constructor(private http: HttpClient) {
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

getApplicationById(application_id: number): Observable<Application> {
  let url = `/api/applications/${application_id}`;
    return new Observable((observer: any) => {
      this.http.get(url)
         .pipe(map((res: any) => res))
         .subscribe((data: Application) => {
            this._application = data

            observer.next(this._application);
            observer.complete();
         });
  });
}

createChatroom(user1_id: string, user2_id: string): Observable<Chatroom> {
  let url = `/api/chatrooms/`;
  return new Observable((observer: any) => {
     this.http.post<any>(url, JSON.stringify({
         "user1_id": user1_id,
         "user2_id": user2_id
     }), {headers: this.headers})
         .pipe(map((res: any) => res))
         .subscribe((data: any) => {
            this._output = data

            observer.next(this._output);
            observer.complete();
         });
  });
}

setSelectedIdx(i: number, post_id: number) {
  this.selectedIdxMap = new Map<number,number>()
  this.selectedIdxMap.set(post_id, i)
}

getSelected(post_id: number) {
  return this.selectedIdxMap.get(post_id)
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
