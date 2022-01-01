import { Injectable } from '@angular/core';
import {Http} from '@angular/http';
import {environment} from '../environments/environment';
import { map } from 'rxjs/operators';

@Injectable()
export class LoginService {

  constructor(private http: Http) { }

  getTitle() {
    return this.http.get(`${environment.serverUrl}/login`).pipe(map((response: any) => {response.json()}));
  }

}