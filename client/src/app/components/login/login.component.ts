import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl } from '@angular/forms'
import { Router } from '@angular/router'

import { HOME } from 'src/app/constants/constants';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {

  public loginError: Boolean;
  public hidePassword: Boolean;
  public loginErrorMessage: string;
  public loginForm: FormGroup;

  constructor(private router: Router) {
    this.hidePassword = true;
    this.loginError = false;
    this.loginErrorMessage = "";
    this.loginForm = new FormGroup({
      username: new FormControl(''),
      password: new FormControl(''),
    });
  }

  ngOnInit(): void { }

  // function to enable the visualization of an error
  public activeErrorMessage(message: string, timeActive: number) {
    this.loginError = true;
    this.loginErrorMessage = message;
    setTimeout(() => this.loginError = false, timeActive);
  }

  public submit(): void {

    if (this.loginForm.valid){
      if ((this.loginForm.controls['username'].value == "dario.ciaudano") && (this.loginForm.controls['password'].value == "coffeearena")){
        this.router.navigate([HOME])
      } else {
        this.activeErrorMessage("Error on credentials!!", 5000)
      }
    }
  }

}
