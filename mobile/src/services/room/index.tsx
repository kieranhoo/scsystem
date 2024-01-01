import axios from "axios";

export const rooms = {
  getgeneralstateallroom: async () => {
    try {
      await axios.get(`${process.env.BASE_URL}/stat/rooms`).then((resp) => {
        console.log(resp.data);
        return resp.data;
      });
    } catch (err: any) {
      return err;
    }
  },

  getbardata: async (room_id: string) => {
    try {
      const result = await axios.get(`${process.env.BASE_URL}/stat/chart`, {
        params: {
          room_id: room_id,
        },
      });
      return result.data;
    } catch (err: any) {
      return err;
    }
  },

  getroominform: async () => {
    try {
      const resp = await axios.get(`${process.env.BASE_URL}/room`);
      return resp.data;
    } catch (err: any) {
      return err;
    }
  },
};
