import axios from "axios"
export interface register {
    "email": string,
    "end_day": string,
    "first_name": string,
    "last_name": string,
    "phone_number": string,
    "registration_time": string,
    "room_id": string,
    "start_day": string,
    "student_id": string,
    "supervisor": string
}
export const lab = {
    register: async () => {
        const paramtest: register = {
            "email": "lamchinh@gmail.com",
            "end_day": "T5",
            "first_name": "Chinh",
            "last_name": "Lam",
            "phone_number": "0123456789",
            "registration_time": "1",
            "room_id": "1",
            "start_day": "1",
            "student_id": "2000000",
            "supervisor": "1"
        }
        let c = await axios.post(`https://api-server-ikierans.cloud.okteto.net/api/v1/lab/register`, paramtest, {
            headers: {
                'Content-Type': 'application/json'
            }
        })
        console.log(c.data)
    }
}
