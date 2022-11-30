import {StudentsInterface} from "./IStudent";
import {SubjectInterface} from "./ISubject";
import {StateInterface} from "./IState";

export interface RegistrationInterface {
    ID?: number,
    Subject?: SubjectInterface,
    SubjectID?: number,
    Student?: StudentsInterface;
    StudentID?: number,
    State?: StateInterface;
    StateID?: number,

}