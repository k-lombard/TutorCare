import { Injectable } from '@angular/core';
import {Http, RequestMethod, RequestOptions} from '@angular/http';
import {environment} from '../../../environments/environment';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { User } from 'src/app/models/user.model';

@Injectable()
export class EditProfileService {
  results:Object[];
  _output: any[] | undefined;
  headers = new HttpHeaders({
    'Content-Type': 'application/json'
});
  constructor(private http: HttpClient) {
    this.results = []
  }

  editProfile(user_id?: string, email?: string, experience?: string, user_category?: string, bio?: string, password?: string, preferences?: string): Observable<User> {
    // let url = `${environment.serverUrl}/api/signup/`;
    let url = `/api/profile/${user_id}`;
    return new Observable((observer: any) => {
       this.http.put<any>(url, JSON.stringify({
           "email": email,
           "user_category": user_category,
           "experience": experience,
           "bio": bio,
           "password": password,
           "preferences": preferences
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
