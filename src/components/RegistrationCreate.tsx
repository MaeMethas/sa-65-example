import React, { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import Button from "@mui/material/Button";
import FormControl from "@mui/material/FormControl";
import Container from "@mui/material/Container";
import Paper from "@mui/material/Paper";
import Grid from "@mui/material/Grid";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import Divider from "@mui/material/Divider";
import Snackbar from "@mui/material/Snackbar";
import Select, { SelectChangeEvent } from "@mui/material/Select";
import MuiAlert, { AlertProps } from "@mui/material/Alert";
import TextField from "@mui/material/TextField";


import { SubjectInterface } from "../interfaces/ISubject";
import { StateInterface } from "../interfaces/IState";
import { RegistrationInterface } from "../interfaces/IRegistration";
import { StudentsInterface } from "../interfaces/IStudent";

import {
    GetRegistration,
    GetState,
    GetSubjects,
    CreateRegistration,
    GetOnlyStudent,
} from "../services/HttpClientService";
const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
    props,
    ref
) {
    return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

function RegistrationCreate() {
    const [student, setStudent] = useState<StudentsInterface>({});
    const [subjects, setSubjects] = useState<SubjectInterface[]>([]);
    const [states, setStates] = useState<StateInterface[]>([]);
    const [registration, setRegistration] = useState<RegistrationInterface>({});
    const [success, setSuccess] = useState(false);
    const [error, setError] = useState(false);

    const handleClose = (
        event?: React.SyntheticEvent | Event,
        reason?: string
    ) => {
        if (reason === "clickaway") {
            return;
        }
        setSuccess(false);
        setError(false);
    };

    const handleChange = (event: SelectChangeEvent) => {
        const name = event.target.name as keyof typeof registration;
        const value = event.target.value;
        setRegistration({
            ...registration,
            [name]: value,
        });
        console.log(`${name}, ${value}`)
    };

    const getStudents = async () => {
        let res = await GetOnlyStudent();
        if (res) {
            setStudent(res);
            registration.StudentID=res.ID
            console.log(res);
        }
    };

    const getSubjects = async () => {
        let res = await GetSubjects();
        if (res) {
            setSubjects(res);
            console.log(res);
        }
    };

    const getStates = async () => {
        let res = await GetState();
        if (res) {
            setStates(res);
            console.log(res);
        }
    };





    useEffect(() => {
        getStates();
        getSubjects();
        getStudents();
    }, []);

    const convertType = (data: string | number | undefined) => {
        let val = typeof data === "string" ? parseInt(data) : data;
        return val;
    };

    async function submit() {
        let data = {
            SubjectID: convertType(registration.SubjectID),
            StudentID: convertType(registration.StudentID),
            StateID: convertType(registration.StateID),
        };
        console.log(data)
        let res = await CreateRegistration(data);
        if (res) {
            setSuccess(true);
        } else {
            setError(true);
        }
    }

    return (
        <Container maxWidth="md">
            <Snackbar
                open={success}
                autoHideDuration={3000}
                onClose={handleClose}
                anchorOrigin={{ vertical: "top", horizontal: "center" }}
            >
                <Alert onClose={handleClose} severity="success">
                    ลงทะเบียนสำเร็จ
                </Alert>
            </Snackbar>
            <Snackbar
                open={error}
                autoHideDuration={6000}
                onClose={handleClose}
                anchorOrigin={{ vertical: "top", horizontal: "center" }}
            >
                <Alert onClose={handleClose} severity="error">
                    ลงทะเบียนไม่สำเร็จ
                </Alert>
            </Snackbar>
            <Paper>
                <Box
                    display="flex"
                    sx={{
                        marginTop: 2,
                    }}
                >
                    <Box sx={{ paddingX: 2, paddingY: 1 }}>
                        <Typography
                            component="h2"
                            variant="h6"
                            color="primary"
                            gutterBottom
                        >
                            ระบบลงทะเบียนเรียน
                        </Typography>
                    </Box>
                </Box>
                <Divider />
                <Grid container spacing={3} sx={{ padding: 2 }}>
                    <Grid item xs={6}>
                        <FormControl fullWidth variant="outlined">
                            <p>รหัสนักศึกษา</p>
                            <Select
                                native

                                value={registration.StudentID + ""}
                                onChange={handleChange}
                                disabled
                                inputProps={{
                                    name: "StudentID",
                                }}
                            >
                                <option aria-label="None" value="">
                                    กรุณารหัสนักศึกษา
                                </option>
                                <option value={student?.ID} key={student?.ID}>
                                    {student?.S_ID}
                                </option>
                            </Select>
                        </FormControl>
                    </Grid>
                    <Grid item xs={8}>
                        <FormControl fullWidth variant="outlined">
                            <p>รายวิชา</p>
                            <Select
                                native
                                value={registration.SubjectID + ""}
                                onChange={handleChange}
                                inputProps={{
                                    name: "SubjectID",
                                }}
                            >
                                <option aria-label="None" value="">
                                    เลือกรายวิชา
                                </option>
                                {subjects.map((item: SubjectInterface) => (
                                    <option value={item.ID} key={item.ID}>
                                        {item.Name}
                                    </option>
                                ))}
                            </Select>
                        </FormControl>
                    </Grid>
                    <Grid item xs={8}>
                        <FormControl fullWidth variant="outlined">
                            <p>สถานะวิชาเรียน</p>
                            <Select
                                native
                                value={registration.StateID + ""}
                                onChange={handleChange}
                                inputProps={{
                                    name: "StateID",
                                }}
                            >
                                <option aria-label="None" value="">
                                    เลือกสถานะการเรียน
                                </option>
                                {states.map((item: StateInterface) => (
                                    <option value={item.ID} key={item.ID}>
                                        {item.Name}
                                    </option>
                                ))}
                            </Select>
                        </FormControl>
                    </Grid>
                    <Grid item xs={12}>
                        <Button
                            component={RouterLink}
                            to="/"
                            variant="contained"
                            color="inherit"
                        >
                            กลับ
                        </Button>
                        <Button
                            style={{ float: "right" }}
                            onClick={submit}
                            variant="contained"
                            color="primary"
                        >
                            บันทึก
                        </Button>
                    </Grid>
                </Grid>
            </Paper>
        </Container>
    );
}

export default RegistrationCreate;