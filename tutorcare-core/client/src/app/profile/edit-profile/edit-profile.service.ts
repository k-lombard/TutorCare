import { Injectable } from '@angular/core';
import {Http, RequestMethod, RequestOptions} from '@angular/http';
import {environment} from '../../../environments/environment';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { User } from 'src/app/models/user.model';
import { Profile } from 'src/app/models/profile.model';

@Injectable()
export class EditProfileService {
  results:Object[];
  _output: any[] | undefined;
  _profile: Profile | undefined;
  headers = new HttpHeaders({
    'Content-Type': 'application/json'
});
  constructor(private http: HttpClient) {
    this.results = []
  }

  editProfile(user_id?: string, updatedData?: Profile): Observable<Profile> {
    // let url = `${environment.serverUrl}/api/signup/`;
    let url = `/api/profile/p/${user_id}`;
    console.log("service")
    console.log(updatedData)
    return new Observable((observer: any) => {
       this.http.put<any>(url, JSON.stringify({
           /*"email": email,
           "user_category": user_category,
           "experience": experience,
           "bio": bio,
           "password": password,
           "preferences": preferences*/
           
          "profile_pic":    updatedData.profile_pic,
          "bio":            updatedData.bio,
          "badge_list":     updatedData.badge_list,
          "age":            updatedData.age,
          "gender":         updatedData.gender,
          "language":       updatedData.language,
          "experience":     updatedData.experience,
          "education":      updatedData.education,
          "skills":         updatedData.skills,
          "service_types":  updatedData.service_types,
          "age_groups":     updatedData.age_groups,
          "covid19":        updatedData.covid19,
          "cpr":            updatedData.cpr,
          "first_aid":      updatedData.first_aid,
          "smoker":         updatedData.smoker,
          "jobs_completed": updatedData.jobs_completed,
          "rate_range":     updatedData.rate_range,
          "rating":         updatedData.rating,
       }), {headers: this.headers})
           .pipe(map((res: any) => res))
           .subscribe((data: any) => {
              this._output = data
              console.log("output")
              console.log(data)
              observer.next(this._output);
              observer.complete();
           });
    });
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
