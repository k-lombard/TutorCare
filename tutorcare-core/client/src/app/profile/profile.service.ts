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


@Injectable()
export class ProfileService {
  results:Object[];
  _chatrooms: Chatroom[] | undefined;
  _output: any[] | undefined;
  _chatroom: Chatroom | undefined
  _messages: Message[] | undefined
  selectedIdx!: number
  _post_id: string | undefined
  headers = new HttpHeaders({
    'Content-Type': 'application/json'
  });
  constructor(private http: HttpClient, private toastr: ToastrService) {
    this.results = []
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



}
