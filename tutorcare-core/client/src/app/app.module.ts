import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { MatInputModule } from '@angular/material/input';
import { MatCardModule } from '@angular/material/card';
import {MatSlideToggleModule} from '@angular/material/slide-toggle';
import {MatCheckboxModule} from '@angular/material/checkbox';
import { NgxSliderModule } from '@angular-slider/ngx-slider';

import { UsersService } from './users.service';
import { HttpModule } from '@angular/http';
import { MatButtonModule } from '@angular/material/button';
import { ReactiveFormsModule, FormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import {MatSidenavModule} from '@angular/material/sidenav';
import {MatIconModule} from '@angular/material/icon';
import {MatListModule} from '@angular/material/list';
import {MatToolbarModule} from '@angular/material/toolbar';
import { SignupService } from './signup/signup.service';
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
import { EditProfileService } from './profile/edit-profile/edit-profile.service';
import { ToastrModule } from 'ngx-toastr';
import { FindJobsService } from './find-jobs/find-jobs.service';
import { ApplyJobService } from './find-jobs/apply-job/apply-job.service';
import { AboutComponent } from './about/about.component';
import { ApplicationsReceivedService } from './find-jobs/applications-received/applications-received.service';
import { ActiveJobsService } from './find-jobs/active-jobs/active-jobs.service';
import { MyJobPostingsService } from './find-jobs/my-job-postings/my-job-postings.service';
import { ChatroomsService } from './find-jobs/chatrooms/chatrooms.service';
import { VerifyService } from './signup/verify/verify.service';
import { SimplebarAngularModule } from 'simplebar-angular';
import { NgSelectModule } from '@ng-select/ng-select';
import { ProfileService } from './profile/profile.service';
import { AppliedToService } from './find-jobs/applied-to/applied-to.service';
import { NgxMatDateAdapter, NgxMatDateFormats, NGX_MAT_DATE_FORMATS } from '@angular-material-components/datetime-picker';
import { NgxMatMomentModule, NgxMatMomentAdapter, NGX_MAT_MOMENT_DATE_ADAPTER_OPTIONS } from '@angular-material-components/moment-adapter';
import { MAT_DATE_LOCALE } from '@angular/material/core';
import {MatTooltipModule} from '@angular/material/tooltip';
import { JobService } from './job/job.service';
// import 'hammerjs';


const CUSTOM_DATE_FORMATS: NgxMatDateFormats = {
  parse: {
    dateInput: 'l, LTS'
  },
  display: {
    dateInput: 'YYYY-MM-DD hh:mm A',
    monthYearLabel: 'MMM YYYY',
    dateA11yLabel: 'LL',
    monthYearA11yLabel: 'MMMM YYYY',
  }
};
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
    NgSelectModule,
    ToastrModule.forRoot(),
    MatInputModule,
    MatCardModule,
    NgxSliderModule,
    MatSlideToggleModule,
    MatCheckboxModule,
    MatTooltipModule,
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
  providers: [UsersService, SignupService, FindCareService, EditProfileService, FindJobsService, ApplyJobService, ApplicationsReceivedService, ActiveJobsService, MyJobPostingsService, ChatroomsService, VerifyService, ProfileService, AppliedToService, JobService, { provide: RouterStateSerializer, useClass: CustomSerializer },
    {
      provide: NgxMatDateAdapter,
      useClass: NgxMatMomentAdapter,
      deps: [MAT_DATE_LOCALE, NGX_MAT_MOMENT_DATE_ADAPTER_OPTIONS]
    },
    { provide: NGX_MAT_DATE_FORMATS, useValue: CUSTOM_DATE_FORMATS }
],
  bootstrap: [AppComponent]
})
export class AppModule { }