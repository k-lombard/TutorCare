import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HomeComponent } from './home/home.component';
import { AboutComponent } from './about/about.component';
import { LoginComponent } from './auth/login/login.component';
import { SignupComponent } from './signup/signup.component';
import { AuthModule } from './auth/auth.module';
import { FindCareComponent } from './find-care/find-care.component';
import { EditProfileComponent } from './profile/edit-profile/edit-profile.component';
import { AuthGuard } from './auth/auth.guard';
import { CreateJobDialog, FindJobsComponent } from './find-jobs/find-jobs.component';
import { ApplicationsReceivedComponent } from './find-jobs/applications-received/applications-received.component';
import { ActiveJobsComponent } from './find-jobs/active-jobs/active-jobs.component';
import { MyJobPostingsComponent } from './find-jobs/my-job-postings/my-job-postings.component';
import { ChatroomsComponent } from './find-jobs/chatrooms/chatrooms.component';
import { VerifyComponent } from './signup/verify/verify.component';
import { ProfileComponent } from './profile/profile.component';
import { AppliedToComponent } from './find-jobs/applied-to/applied-to.component';
import { JobComponent } from './job/job.component';
import { SettingsComponent } from './settings/settings.component';
import { PastJobsComponent } from './settings/past-jobs/past-jobs.component';

const routes: Routes = [
  { path: 'home', component: HomeComponent },
  { path: 'login', component: LoginComponent },
  { path: 'signup', component: SignupComponent },
  { path: 'find-care', component: FindCareComponent },
  { path: 'find-jobs', component: FindJobsComponent },
  { path: 'about-us', component: AboutComponent },
  { path: 'profile/:id', component: ProfileComponent, canActivate: [AuthGuard]},
  { path: 'profile/:id/edit-profile', component: EditProfileComponent, canActivate: [AuthGuard]},
  { path: 'find-jobs/create', component: CreateJobDialog, canActivate: [AuthGuard]},
  { path: 'find-jobs/applications-received', component: ApplicationsReceivedComponent, canActivate: [AuthGuard]},
  { path: 'find-jobs/applications-received/:id', component: ApplicationsReceivedComponent, canActivate: [AuthGuard]},
  { path: 'find-jobs/active-jobs', component: ActiveJobsComponent, canActivate: [AuthGuard]},
  { path: 'find-jobs/my-job-postings', component: MyJobPostingsComponent, canActivate: [AuthGuard]},
  { path: 'find-jobs/jobs/:id', component: JobComponent, canActivate: [AuthGuard]},
  { path: 'find-jobs/applied-to', component: AppliedToComponent, canActivate: [AuthGuard]},
  { path: 'find-jobs/applied-to/:id', component: AppliedToComponent, canActivate: [AuthGuard]},
  { path: 'find-jobs/messages', component: ChatroomsComponent, canActivate: [AuthGuard]},
  { path: 'find-jobs/messages/:id', component: ChatroomsComponent, canActivate: [AuthGuard]},
  { path: 'verify', component: VerifyComponent},
  { path: 'settings', component: SettingsComponent, canActivate: [AuthGuard]},
  { path: 'settings/past-jobs', component: PastJobsComponent, canActivate: [AuthGuard]},
  { path: '',   redirectTo: '/home', pathMatch: 'full' },
  { path: '**', redirectTo: '/home', pathMatch: 'full' },
  { path: '',   component: HomeComponent }

];

@NgModule({
  imports: [RouterModule.forRoot(routes), AuthModule],
  exports: [RouterModule]
})
export class AppRoutingModule { }
