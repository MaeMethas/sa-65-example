import React, { useState, useEffect } from "react";
import Typography from "@mui/material/Typography";
import Container from "@mui/material/Container";
import Box from "@mui/material/Box";
import { DataGrid, GridColDef } from "@mui/x-data-grid";
import { RegistrationInterface } from "../interfaces/IRegistration";
import { GetRegistration } from "../services/HttpClientService";

function Registration() {
    const [Registration, setRegistration] = useState<RegistrationInterface[]>([]);
  
    useEffect(() => {
      getRegistration();
    }, []);
  
    const getRegistration = async () => {
      let res = await GetRegistration();
      if (res) {
        console.log("res");
        console.log(res);
        setRegistration(res);
      } 
    };

    const columns: GridColDef[] = [
        { field: "ID", headerName: "ลำดับ", width: 50 },
        {
            field: "Student",
            headerName: "รหัสนักศึกษา",
            width: 200,
            valueFormatter: (params) => params.value.Name,
        },
        {
            field: "Subject",
            headerName: "รหัสวิชา",
            width: 200,
            valueFormatter: (params) => params.value.CODE,
        },
        {
            field: "State",
            headerName: "สถานะการเรียน",
            width: 150,
            valueFormatter: (params) => params.value.Name,
        },

    ];

    return (
        <div>
          <Container maxWidth="md">
            <Box
              display="flex"
              sx={{
                marginTop: 2,
              }}
            >
              <Box flexGrow={1}>
                <Typography
                  component="h2"
                  variant="h6"
                  color="primary"
                  gutterBottom
                >
                  ข้อมูลการลงทะเบียนเรียน
                </Typography>
              </Box>
              
            </Box>
            <div style={{ height: 400, width: "100%", marginTop: "20px" }}>
              <DataGrid
                rows={Registration}
                getRowId={(row) => row.ID}
                columns={columns}
                pageSize={5}
                rowsPerPageOptions={[5]}
              />
            </div>
          </Container>
        </div>
      );
}

export default Registration;