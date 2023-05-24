import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { SERVER_HOST } from 'src/app/constants/constants';

@Component({
  selector: 'app-quiz-room',
  templateUrl: './quiz-room.component.html',
  styleUrls: ['./quiz-room.component.scss']
})
export class QuizRoomComponent implements OnInit {

  public currentQuestionIndex = 0;
  public userAnswers = [];
  public questions = [

  ];

  constructor(private router: Router, private http: HttpClient) { }

  ngOnInit(): void {
    this.http.get<any>(SERVER_HOST+"qaList").subscribe(
      (response) => {
        let questions = response["questions"];
        this.questions = questions;
      },
      (error) => {
        console.log(error);
      }
    )
  }

  public chooseQuestion(question_index: number){
    this.currentQuestionIndex = question_index;
  }

  public selectAnswer(answer: any){

  }

  public submit(){

  }

}
