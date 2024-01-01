import axios from "axios"

export const activity = {
    checkin: async (registration_id: string) => {
        try {
            const result = await axios.post(`${process.env.BASE_URL}/room/activity`,
                {
                    activity_type: "in",
                    admin_id: "1",
                    registration_id: registration_id
                }
            )
            return true
        }
        catch (err: any) {
            throw err
        }
    },
    checkout: async (registration_id: string) => {
        try {
            const result = await axios.post(`${process.env.BASE_URL}/room/activity`,
                {
                    activity_type: "out",
                    admin_id: "1",
                    registration_id: registration_id
                }
            )
            return true
        }
        catch (err: any) {
            throw err
        }
    }
}

export const room = {
    check: async (sid: string, room_id: string) => {
        try {
            const result = await axios.get(`${process.env.BASE_URL}/room/activity`, {
                params: {
                    sid: sid,
                    room: room_id,
                }
            })
            return result.data
        }
        catch (err: any) {
            throw err
        }
    }
}
