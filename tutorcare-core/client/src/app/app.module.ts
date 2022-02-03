import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { MatInputModule } from '@angular/material/input';
import { MatCardModule } from '@angular/material/card';

import { UsersService } from './users.service';
import { Http, ConnectionBackend, HttpModule } from '@angular/http';
import { MatButtonModule } from '@angular/material/button';
import { ReactiveFormsModule, FormsModule } from '@angular/forms';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import {MatSidenavModule} from '@angular/material/sidenav';
import {MatIconModule} from '@angular/material/icon';
import {MatListModule} from '@angular/material/list';
import {MatToolbarModule} from '@angular/material/toolbar';
import { LoginComponent } from './auth/login/login.component';
import { NavbarComponent } from './navbar/navbar.component';
import { SignupComponent } from './signup/signup.component';
import { SignupService } from './signup/signup.service';
import { HomeComponent } from './home/home.component';
import {MatGridListModule} from '@angular/material/grid-list';
import { StoreModule } from '@ngrx/store';
import { AuthModule } from './auth/auth.module';
import { reducers, metaReducers } from './reducers';
import { environment } from 'src/environments/environment';
import { EffectsModule } from '@ngrx/effects';
import { StoreRouterConnectingModule, RouterStateSerializer } from '@ngrx/router-store';
import { StoreDevtoolsModule } from '@ngrx/store-devtools';
import { CustomSerializer } from './shared/utils';
import { RouterModule } from '@angular/router';
import { FindCareService } from './find-care/find-care.service';
import { BarRatingModule } from 'ngx-bar-rating';
import { EditProfileService } from './account/edit-profile/edit-profile.service';
import { ToastrModule } from 'ngx-toastr';
import { AuthInterceptor } from './auth/auth.interceptor';
import { FindJobsService } from './find-jobs/find-jobs.service';
import { ApplyJobService } from './find-jobs/apply-job/apply-job.service';
import { AboutComponent } from './about/about.component';
import { ApplicationsReceivedService } from './find-jobs/applications-received/applications-received.service';
import { ActiveJobsService } from './find-jobs/active-jobs/active-jobs.service';
import { MyJobPostingsService } from './find-jobs/my-job-postings/my-job-postings.service';
import { ChatroomsService } from './find-jobs/chatrooms/chatrooms.service';
import { VerifyService } from './signup/verify/verify.service';
import { SimplebarAngularModule } from 'simplebar-angular';
import { ThemeService } from './theme.service';

@NgModule({
  declarations: [
    AppComponent,
    AboutComponent
  ],
  imports: [
    RouterModule,
    FormsModule,
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    ToastrModule.forRoot(),
    MatInputModule,
    MatCardModule,
    HttpClientModule,
    MatButtonModule,
    BarRatingModule,
    ReactiveFormsModule,
    HttpModule,
    MatToolbarModule,
    MatSidenavModule,
    SimplebarAngularModule,
    MatListModule,
    MatIconModule,
    MatGridListModule,
    AuthModule.forRoot(),
    StoreModule.forRoot(reducers, { metaReducers }),
    !environment.production ? StoreDevtoolsModule.instrument() : [],
    EffectsModule.forRoot([]),
    StoreRouterConnectingModule.forRoot({stateKey:'router'})
  ],
  providers: [UsersService, SignupService, FindCareService, EditProfileService, FindJobsService, ApplyJobService, ApplicationsReceivedService, ActiveJobsService, MyJobPostingsService, ChatroomsService, VerifyService, ThemeService, { provide: RouterStateSerializer, useClass: CustomSerializer },],
  bootstrap: [AppComponent]
})
export class AppModule { }
