import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/Observable';
import { map } from 'rxjs/operators';


import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Application } from 'src/app/models/application.model';


@Injectable()
export class ApplyJobService {
  results:Object[];
  _output: any[] | undefined;
  headers = new HttpHeaders({
    'Content-Type': 'application/json'
  });
  constructor(private http: HttpClient) {
    this.results = []
  }


 applyJob(userId: string, postId: number, message: string, poster: string): Observable<Application> {
  if (poster !== userId) {
    let url = `/api/applications/`;
    return new Observable((observer: any) => {
      // @ts-ignore
      this.http.post<any>(url, JSON.stringify({
          "user_id": userId,
          "post_id": postId,
          "message": message
      }), {headers: this.headers})
          .pipe(map((res: any) => res))
          .subscribe((data: any) => {
              this._output = data

              observer.next(this._output);
              observer.complete();

          });
    });
  } else {
    return undefined
  }
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
