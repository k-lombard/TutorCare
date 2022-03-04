import { Injectable } from '@angular/core';
import { Observable, throwError } from 'rxjs';
import { catchError, map } from 'rxjs/operators';
import { HttpClient, HttpErrorResponse, HttpHeaders } from '@angular/common/http';
import { ToastrService } from 'ngx-toastr';
import { User } from '../models/user.model';
import { Profile } from '../models/profile.model';

@Injectable()
export class AccountService {
  results: Object[];
  _profile: Profile | undefined;
  headers = new HttpHeaders({
    'Content-Type': 'application/json'
  });
  constructor(private http: HttpClient, private toastr: ToastrService) {
    this.results = []
  }

  getProfile(user_id: string): Observable<Profile> {
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

}
