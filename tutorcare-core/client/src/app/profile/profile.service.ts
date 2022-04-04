import { Injectable } from '@angular/core';
import {Http} from '@angular/http';
import {environment} from '../../environments/environment';
import { Observable } from 'rxjs/Observable';
import { catchError, map } from 'rxjs/operators';
import { GeolocationPositionWithUser } from '../models/geolocationposition.model';
import { Post } from '../models/post.model';
import { HttpClient, HttpErrorResponse, HttpHeaders } from '@angular/common/http';
import { Application } from 'src/app/models/application.model';
import { throwError } from 'rxjs';
import { ToastrService } from 'ngx-toastr';
import { Chatroom } from 'src/app/models/chatroom.model';
import { Message } from 'src/app/models/message.model';
import { ThisReceiver } from '@angular/compiler';
import { User } from '../models/user.model';
import { Profile } from '../models/profile.model';
import { Review } from '../models/review.model';


@Injectable()
export class ProfileService {
  results:Object[];
  _output: any[] | undefined;
  _post_id: string | undefined
  _profile: Profile | undefined;
  headers = new HttpHeaders({
    'Content-Type': 'application/json'
  });
  constructor(private http: HttpClient, private toastr: ToastrService) {
    this.results = []
  }

  getProfileByUserId(user_id: string): Observable<Profile> {
    let url = `/api/profile/p/${user_id}`;
    return new Observable((observer: any) => {
       this.http.get(url)
           .pipe(map((res: any) => res))
           .subscribe((data: Profile) => {
              this._profile = data
              console.log(data)
              observer.next(this._profile);
              observer.complete();
           });
    });
 }

  getUserByUserId(user_id: string): Observable<User> {
    let url = `/api/users/${user_id}`;
    return new Observable((observer: any) => {
       this.http.get(url)
           .pipe(map((res: any) => res))
           .subscribe((data: User) => {

              observer.next(data);
              observer.complete();
           });
    });
 }

 getReviewsByUserId(user_id: string): Observable<Review[]> {
   let url = `/api/reviews/${user_id}`;
   return new Observable((observer: any) => {
      this.http.get(url)
          .pipe(map((res: any) => res.reviews))
          .subscribe((data: Review[]) => {
             observer.next(data);
             observer.complete();
          });
   });
 }



}
