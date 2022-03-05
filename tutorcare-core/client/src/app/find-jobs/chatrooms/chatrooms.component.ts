import { ThisReceiver } from '@angular/compiler';
import {Component, OnInit, ChangeDetectionStrategy, ViewChild, ElementRef, AfterViewInit, OnChanges, Directive} from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';
import { ActivatedRoute, NavigationEnd, Router } from '@angular/router';
import { select, Store } from '@ngrx/store';
import { DefaultGlobalConfig, ToastrService } from 'ngx-toastr';
import { BehaviorSubject, Subject, Subscription } from 'rxjs';
import { map, switchMap } from 'rxjs/operators';
import SimpleBar from 'simplebar';
import { SimplebarAngularComponent } from 'simplebar-angular';
import { getCurrUser } from 'src/app/auth/auth.selectors';
import { Application } from 'src/app/models/application.model';
import { Chatroom } from 'src/app/models/chatroom.model';
import { Message } from 'src/app/models/message.model';
import { Post } from 'src/app/models/post.model';
import { User } from 'src/app/models/user.model';
import { AppState } from 'src/app/reducers';
import { GeolocationPositionWithUser } from '../../models/geolocationposition.model';
import { ChatroomsService } from './chatrooms.service';

@Component({
  selector: 'chatrooms-component',
  templateUrl: './chatrooms.component.html',
  styleUrls: ['./chatrooms.component.scss']
})
export class ChatroomsComponent implements OnInit{

		public acceptSubject: Subject<boolean> = new BehaviorSubject<boolean>(false);
    public acceptActive = this.acceptSubject.asObservable();
    selectedValue: string | undefined
    userCategory: string = ""
    user!: User
    userId!: string
    chatrooms!: Chatroom[]
    user1!: User
    user2: User
    user1_id!: string
    user2_id!: string
    currChatroom!: Chatroom
    locs!: GeolocationPositionWithUser[]
    mySubscription!: any
    chatroomId!: number
    userType!: string
    ws!: WebSocket
    editable: boolean = true
    messages!: Message[]
    messageForm: FormGroup
    otherUser: User
    otherUserId: string
    menuVisible: boolean
    chatToken!: string
    url: string = ""
    options: any = {classNames: {
      // defaults
      content: 'simplebar-content',
      scrollContent: 'simplebar-scroll-content',
      scrollbar: 'simplebar-scrollbar',
      track: 'simplebar-track'
    }}
    private routeSub: Subscription;
    constructor(private router: Router, private chatroomService: ChatroomsService, private store: Store<AppState>, private route: ActivatedRoute, private toastr: ToastrService, private fb: FormBuilder, public elementRef: ElementRef) {}
    @ViewChild('scrollMe') public myScrollContainer: ElementRef;
    async ngOnInit() {

      this.messageForm = new FormGroup({
        message: new FormControl('')
      })
      this.store
      .pipe(
          select(getCurrUser)
      ).subscribe(async data =>  {
          this.user = data
          this.userId = this.user.user_id || ""
          this.userType = this.user.user_category
          //console.log(this.userId)
          //console.log(this.chatToken)
      })
      this.routeSub = this.route.params.subscribe(params => {
        this.chatroomId = parseInt(params['id'])
        //console.log(this.chatroomId)
      });

      this.chatroomService.getChatroomsByUserId(this.userId).subscribe(data => {
        this.chatrooms = data
        if (this.chatrooms) {
          for (let i = 0; i < this.chatrooms?.length; i++) {
            if (this.chatrooms[i].chatroom_id === this.chatroomId) {
              this.chatroomService.setSelectedIdx(i)
            }
          }
        }
      })
      if (this.chatroomId) {
        this.chatroomService.getChatroomById(this.chatroomId).subscribe(chatroom => {
          this.currChatroom = chatroom
          this.user1 = chatroom.user1
          this.user2 = chatroom.user2
          this.user1_id = chatroom.user1_id
          this.user2_id = chatroom.user2_id
          if (this.user1_id == this.userId) {
            this.otherUser = this.user2
            this.otherUserId = this.user2_id
          } else {
            this.otherUser = this.user1
            this.otherUserId = this.user1_id
          }
        })
        if (!this.messages || this.messages.length === 0) {
          this.chatroomService.getMessagesByChatroomId(this.chatroomId).subscribe(messages => {
            console.log(messages)
            this.messages = messages
          })
        }
        // this.scrollBottom()
      }
      if (this.chatroomId) {
        this.chatroomService.getChatroomToken(this.userId).then((res: string) => {
          this.url = "ws://" + "localhost:8080" + "/api/" + this.chatroomId + "/ws";
          console.log(this.url)
          this.ws = new WebSocket(this.url)
          console.log(this.ws)
          this.ws.onmessage = (msg) => {
            const socket = msg.target as WebSocket;
            if (socket.url !== this.url) return socket.close();
            const d = JSON.parse(msg.data);
            this.messages = [...this.messages, d];
          }
          this.ws.onopen = () => {
            console.log('websocket connected!');
            this.ws.send(JSON.stringify({
              "user_id": this.userId,
              "token": res
            }))
          }
          // this.ws.onclose = () => console.log("websocket closed");
          // window.addEventListener('beforeunload', () => this.ws.close())
        })
      }
  }

  createMessage(sender: User, message: string, timestamp: string) {
    return {
      sender: sender,
      message: message,
      timestamp: timestamp
    }
  }

  ngOnDestroy() {
    this.routeSub.unsubscribe();
    // if (this.ws) {
    //   this.ws.close()
    // }
    // return () => {
    //   window.removeEventListener('beforeunload', () => this.ws.close())
    // }
  }

    sendMessage() {
      const sender = this.user
      const id = this.userId
      const msg = this.messageForm.get('message').value
      if (msg.length > 0) {
        var msgStr: string  = JSON.stringify({
          message: msg,
          sender_id: id,
          chatroom_id: this.chatroomId,
          sender: { first_name: sender.first_name, last_name: sender.last_name },
          timestamp: "Just now"
        })
        this.ws.send(msgStr)
        this.chatroomService.sendMessage(this.messageForm.get('message').value, this.userId, this.currChatroom.chatroom_id).subscribe(message => {
          console.log(message)
          this.messageForm.reset()
        })
      }
    }

    scrollBottom(){
      this.myScrollContainer['SimpleBar'].getScrollElement().scrollTop = this.myScrollContainer['SimpleBar'].getScrollElement().scrollHeight
    }

    onFindCareClick() {
        this.router.navigate(['/find-care'])
    }

    setChatroom(chatroom: Chatroom) {
      this.currChatroom = chatroom
      this.user1 = chatroom.user1
      this.user2 = chatroom.user2
      this.user1_id = chatroom.user1_id
      this.user2_id = chatroom.user2_id
      this.chatroomId = this.currChatroom.chatroom_id
      if (this.user1_id === this.userId) {
        this.otherUser = this.user2
        this.otherUserId = this.user2_id
      } else {
        this.otherUser = this.user1
        this.otherUserId = this.user1_id
      }
      this.chatroomService.getMessagesByChatroomId(this.currChatroom.chatroom_id).subscribe(messages => {
        console.log(messages)
        this.messages = messages
      })
    }

    back() {
      this.currChatroom = undefined
    }

    backToMenu() {
      this.currChatroom = undefined
      this.menuVisible = true
    }

    onEditClick() {
      this.editable = false
      console.log(this.editable)
    }

    setSelected(i: number) {
      this.chatroomService.setSelectedIdx(i)
    }

    getSelected() {
      return this.chatroomService.getSelected()
    }






}
