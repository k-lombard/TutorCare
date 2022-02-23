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
import { Chatroom } from 'src/app/models/chatroom.model';
import { Message } from 'src/app/models/message.model';
import { ThisReceiver } from '@angular/compiler';
import { ApplicationsReceivedService } from '../applications-received/applications-received.service';
import { Router } from '@angular/router';
import { ObserversModule } from '@angular/cdk/observers';


@Injectable()
export class ChatroomsService {
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
  constructor(private http: HttpClient, private toastr: ToastrService, private applicationsService: ApplicationsReceivedService, private router: Router) {
    this.results = []
  }

  getChatroomsByUserId(user_id: string): Observable<Chatroom[]> {
    let url = `/api/chatrooms/user/${user_id}`;
    return new Observable((observer: any) => {
       this.http.get(url)
           .pipe(map((res: any) => res.chatrooms))
           .subscribe((data: Chatroom[]) => {
              this._chatrooms = data

              observer.next(this._chatrooms);
              observer.complete();
           });
    });
 }

 getMessagesByChatroomId(chatroom_id: number): Observable<Message[]> {
  let url = `/api/messages/chatroom/${chatroom_id}`;
  return new Observable((observer: any) => {
     this.http.get(url)
         .pipe(map((res: any) => res.messages))
         .subscribe((data: Message[]) => {
            this._messages = data

            observer.next(this._messages);
            observer.complete();
         });
  });
 }

 sendMessage(message: string, sender_id: string, chatroom_id: number): Observable<Message> {
  let url = `/api/messages/`;
  return new Observable((observer: any) => {
     this.http.post<any>(url, JSON.stringify({
         "sender_id": sender_id,
         "message": message,
         "chatroom_id": chatroom_id
     }), {headers: this.headers})
         .pipe(map((res: any) => res))
         .subscribe((data: any) => {
            this._output = data
            observer.next(this._output);
            observer.complete();
         });
  });
}

getChatroomByTwoUsers(user1_id: string, user2_id: string) {
  let url = `/api/chatrooms/users/${user1_id}/${user2_id}`;
  if (user1_id == user2_id) {
    this.toastr.error("Error: Cannot create a chatroom between a single user")
    return undefined
  }
  return new Observable((observer: any) => {
     this.http.get(url)
         .pipe(map((res: any) => res),
         catchError((err: HttpErrorResponse) => {
           this.createChatroom(user1_id, user2_id).subscribe((chat: Chatroom) => {
            this.toastr.success("Success: New chatroom created with ID: " + chat.chatroom_id, "Success", {closeButton: true, timeOut: 5000, progressBar: true});
            this.router.navigate([`/find-jobs/messages/${chat.chatroom_id}`])
          })
          return throwError(err)
        })
         )
         .subscribe((data: Chatroom) => {
            this._chatroom = data

            observer.next(this._chatroom);
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
         .pipe(map((res: any) => res),
         catchError((err: HttpErrorResponse) => {
          return throwError(err)
        })
         )
         .subscribe((data: any) => {
            this._output = data

            observer.next(this._output);
            observer.complete();
         });
  });
}

async getChatroomToken(user_id: string): Promise<string> {
  let url = `api/chatrooms/websocket/${user_id}`
  return await this.http.get<string>(url, {headers: this.headers}).toPromise()
}

setSelected(chatRoomId: number) {
  for (let i = 0; i < this._chatrooms?.length; i++) {
    if (this._chatrooms[i]?.chatroom_id === chatRoomId) {
      this.selectedIdx = i;
    }
  }
}


setSelectedIdx(i: number) {
  this.selectedIdx = i
}

getSelected() {
  return this.selectedIdx
}

getChatroomById(chatroom_id: number): Observable<Chatroom> {
  let url = `/api/chatrooms/${chatroom_id}`;
    return new Observable((observer: any) => {
      this.http.get(url)
         .pipe(map((res: any) => res))
         .subscribe((data: Chatroom) => {
            this._chatroom = data

            observer.next(this._chatroom);
            observer.complete();
         });
  });
}



}
