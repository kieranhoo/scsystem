import axios from "axios";

export const history = {
  getroominform: async () => {
    try {
      const resp = await axios.get(`${process.env.BASE_URL}/room`);
      // console.log(resp.data);
      return resp.data;
    } catch (err: any) {
      return err;
    }
  },
};
