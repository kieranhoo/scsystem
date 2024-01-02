import React, { useState, useEffect } from "react";
import {
  View,
  Text,
  StyleSheet,
  Dimensions,
  TouchableOpacity,
  FlatList,
  Modal,
  TouchableWithoutFeedback,
} from "react-native";
import { RoomsHeader } from "@/components/header/rooms-header";
import CalendarStrip from "react-native-calendar-strip";
import moment, { Moment } from "moment";
import Ionicons from "react-native-vector-icons/Ionicons";
import Octicons from "react-native-vector-icons/Octicons";
import { Colors } from "../../theme/variables";
import axios from "axios";

const screenWidth = Dimensions.get("screen").width;
const screenHeight = Dimensions.get("screen").height;

interface Item {
  activity_type: string;
  admin_id: string;
  email: string;
  end_day: string;
  first_name: string;
  id: any;
  last_name: string;
  office_name: string;
  org_name: string;
  phone_number: string;
  registration_time: string;
  room_name: string;
  start_day: string;
  supervisor: string;
  time: string;
  user_id: string;
}
interface InputPopupProps {
  isVisible: boolean;
  onClose: () => void;
  selectedItem: Item | null;
}

const InputPopup: React.FC<InputPopupProps> = ({ isVisible, onClose, selectedItem }) => {
  const formattedDateTime = selectedItem?.time
  ? moment(selectedItem.time).format('DD/MM/YYYY - HH:mm:ss')
  : '';
  // const [value, setValue] = useState("");
  const handleClick = () => {
    onClose();
  };
  return (
    <Modal
      animationType="none"
      transparent={true}
      visible={isVisible}
      onRequestClose={onClose}
    >
      <TouchableWithoutFeedback onPress={onClose}>
        <View style={styles.modalContainer}>
          <View style={styles.modalContent}>
            <Text style={styles.label}>{selectedItem?.last_name} {selectedItem?.first_name}</Text>
            <Text style={styles.content}>{selectedItem?.user_id}</Text>
            <Text style={styles.content}>{selectedItem?.email}</Text>
            <Text style={styles.inform}>Room: {selectedItem?.room_name}</Text>
            <Text style={styles.inform}>Supervisor: {selectedItem?.supervisor}</Text>
            <View style={styles.horizontalLine}></View>
            <Text style={styles.inform}>{selectedItem?.activity_type === 'in' ? 'Check in: ' : 'Check out: '} {formattedDateTime}</Text>
            <TouchableOpacity onPress={handleClick} style={styles.button}>
              <Text style={styles.buttontext}>OK</Text>
            </TouchableOpacity>
          </View>
        </View>
      </TouchableWithoutFeedback>
    </Modal>
  );
};

export const Rooms = () => {
  const [selectedDate, setSelectedDate] = useState<Moment>(moment());
  const [selectedRoomId, setSelectedRoomId] = useState<string>("");
  const [isModalVisible, setIsModalVisible] = useState(false);
  const [events, setEvents] = useState<Item[]>([]);
  const [totalIn, setTotalIn] = useState<number>(0);
  const [totalOut, setTotalOut] = useState<number>(0);
  const [modalItem, setModalItem] = useState<Item | null>(null);

  const handleDateSelected = (date: Moment) => {
    setSelectedDate(date);
    if (selectedRoomId) {
      fetchDataByRoomAndDate(selectedRoomId, date.format("YYYY-MM-DD"));
    }
  };

  const getDataByRoomAndDate = async (room_id: string, date: string) => {
    try {
      const result = await axios.get(`${process.env.BASE_URL}/room/history`, {
        params: {
          room_id: room_id,
          date: date,
        },
      });
      return result.data;
    } catch (err: any) {
      return [];
    }
  };

  const fetchDataByRoomAndDate = async (room_id: string, date: string) => {
    try {
      const fetchdata = await getDataByRoomAndDate(room_id, date);
      if (fetchdata && fetchdata.data) {
        const data = fetchdata.data.map((item: Item) => ({
          activity_type: item.activity_type,
          admin_id: item.admin_id,
          email: item.email,
          end_day: item.end_day,
          first_name: item.first_name,
          id: item.id,
          last_name: item.last_name,
          office_name: item.office_name,
          org_name: item.org_name,
          phone_number: item.phone_number,
          registration_time: item.registration_time,
          room_name: item.room_name,
          start_day: item.start_day,
          supervisor: item.supervisor,
          time: item.time,
          user_id: item.user_id,
        }));
        setEvents(data);
        setTotalIn(fetchdata.total_in);
        setTotalOut(fetchdata.total_out);
      } else {
        setEvents([]);
        setTotalIn(0);
        setTotalOut(0);
      }
    } catch (error) {
      console.error("Error fetching data:", error);
    }
  };

  const handleCheckStudentPress = (selectedItem: Item) => {
    setIsModalVisible(true);
    setModalItem(selectedItem);
  };

  useEffect(() => {
    if (selectedRoomId && selectedDate) {
      fetchDataByRoomAndDate(selectedRoomId, selectedDate.format("YYYY-MM-DD"));
    }
  }, [selectedRoomId, selectedDate]);

  const renderItem = ({ item }: { item: Item }) => {
    const formattedTime = moment(item.time).format("HH : mm");
    return (
      <TouchableOpacity
        style={styles.informContainer}
        onPress={() => handleCheckStudentPress(item)}
      >
        <View style={styles.informTextContainer}>
          <View style={styles.textNameContainer}>
            <Text style={styles.textName}>
              {item.last_name} {item.first_name}
            </Text>
          </View>
          <View style={styles.textIdContainer}>
            <Text style={styles.textId}>MSSV: {item.user_id}</Text>
          </View>
        </View>
        <View style={styles.informStatusContainer}>
          <View
            style={[
              styles.textStatusContainer,
              {
                backgroundColor:
                  item.activity_type === "out"
                    ? "rgba(237, 74, 74, 0.15)"
                    : "rgba(52, 168, 83, 0.15)",
                borderColor:
                  item.activity_type === "out"
                    ? "rgba(237, 74, 74, 0.87)"
                    : "rgba(52, 168, 83, 0.89)",
              },
            ]}
          >
            <Text
              style={[
                styles.textStatus,
                {
                  color:
                    item.activity_type === "out"
                      ? "rgba(237, 74, 74, 0.87)"
                      : "rgba(52, 168, 83, 0.89)",
                },
              ]}
            >
              Check {item.activity_type}
            </Text>
          </View>
          <View style={styles.textTimeContainer}>
            <Text style={styles.textTime}>{formattedTime}</Text>
          </View>
        </View>
      </TouchableOpacity>
    );
  };

  return (
    <View style={styles.container}>
      <RoomsHeader
        onSelectRoom={(roomId) => setSelectedRoomId(roomId)}
        supervisor={events.length > 0 ? events[0].supervisor : ""}
      />

      <View style={styles.content_container}>
        <View style={styles.calendarContainer}>
          <CalendarStrip
            style={styles.calendarStripStyle}
            showMonth={false}
            scrollable
            scrollerPaging
            numDaysInWeek={7}
            dateNumberStyle={{ color: "#6C6C6C" }}
            dateNameStyle={{ color: "#6C6C6C" }}
            dayContainerStyle={styles.dateContainer}
            highlightDateContainerStyle={styles.highlightDateContainer}
            highlightDateNameStyle={{ color: "#1B61B5" }}
            highlightDateNumberStyle={{ color: "#1B61B5" }}
            selectedDate={selectedDate}
            onDateSelected={handleDateSelected}
          />
        </View>
        <View style={styles.statusContainer}>
          <View style={styles.statusStyle}>
            <View style={styles.checkIn}>
              <View style={styles.checkNum}>
                <Octicons
                  name="people"
                  size={30}
                  color="rgba(52, 168, 83, 0.93)"
                />
                <Text style={styles.regular25In}>{totalIn}</Text>
              </View>
              <Text style={styles.regular14}>check in</Text>
            </View>
            <View style={styles.verticalLine}></View>
            <View style={styles.checkOut}>
              <View style={styles.checkNum}>
                <Octicons
                  name="people"
                  size={30}
                  color="rgba(237, 74, 74, 0.87)"
                />
                <Text style={styles.regular25Out}>{totalOut}</Text>
              </View>
              <Text style={styles.regular14}>check out</Text>
            </View>
          </View>
          {/* <TouchableOpacity
            style={styles.addRecordStyle}
            onPress={handleAddRecordPress}
          >
            <Text style={styles.regular14Add}>Add record</Text>
            <Ionicons
              name="add-circle-outline"
              size={25}
              color="rgba(27, 97, 181, 0.89)"
            />
          </TouchableOpacity> */}
        </View>
        <View style={styles.detailContainer}>
          <FlatList
            data={events}
            renderItem={renderItem}
            keyExtractor={(item) => item.id.toString()}
            showsVerticalScrollIndicator={false}
          />
        </View>
      </View>
      {isModalVisible && (
        <InputPopup
          isVisible={isModalVisible}
          onClose={() => setIsModalVisible(false)}
          selectedItem={modalItem}
        />
      )}
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: "#F4F7FB",
  },
  content_container: {
    flex: 1,
  },
  calendarContainer: {
    flex: 0.12,
    justifyContent: "center",
    marginBottom: 10,
  },
  calendarStripStyle: {
    height: 75,
  },
  dateContainer: {
    borderWidth: 1.5,
    borderRadius: 10,
    borderColor: Colors.WHITE,
    backgroundColor: Colors.WHITE,
  },
  highlightDateContainer: {
    borderColor: "rgba(27, 97, 181, 0.89)",
    backgroundColor: "rgba(27, 97, 181, 0.1)",
  },
  statusContainer: {
    flex: 0.19,
    marginHorizontal: screenWidth * 0.05,
  },
  statusStyle: {
    flex: 0.65,
    flexDirection: "row",
    borderRadius: 10,
    justifyContent: "center",
    backgroundColor: Colors.WHITE,
    shadowOffset: {
      width: 0,
      height: 0,
    },
    shadowOpacity: 0.2,
    shadowRadius: 5,
  },
  checkIn: {
    flex: 0.5,
    alignItems: "center",
    justifyContent: "center",
  },
  verticalLine: {
    height: "80%",
    width: 1,
    alignSelf: "center",
    backgroundColor: "rgba(108, 108, 108, 0.89)",
  },
  checkOut: {
    flex: 0.5,
    alignItems: "center",
    justifyContent: "center",
  },
  checkNum: {
    flex: 0.55,
    flexDirection: "row",
  },
  addRecordStyle: {
    flex: 0.35,
    alignItems: "center",
    justifyContent: "center",
    flexDirection: "row",
    marginTop: screenHeight * 0.005,
    borderBottomLeftRadius: 10,
    borderBottomRightRadius: 10,
    backgroundColor: Colors.WHITE,
    shadowOffset: {
      width: 0,
      height: 0,
    },
    shadowOpacity: 0.1,
    shadowRadius: 2.5,
  },
  detailContainer: {
    flex: 0.69,
    marginBottom: 10,
    marginHorizontal: screenWidth * 0.05,
  },
  informContainer: {
    flex: 1,
    flexDirection: "row",
    padding: 10,
    marginBottom: 10,
    borderRadius: 10,
    backgroundColor: Colors.WHITE,
    shadowOffset: {
      width: 0,
      height: 0,
    },
    shadowOpacity: 0.1,
    shadowRadius: 2.5,
  },
  informTextContainer: {
    flex: 0.6,
  },
  textNameContainer: {
    flex: 0.5,
    paddingLeft: 10,
  },
  textIdContainer: {
    flex: 0.5,
    paddingLeft: 10,
  },
  textName: {
    fontSize: 20,
    fontFamily: "Poppins_400Regular",
    color: Colors.BLACK,
  },
  textId: {
    fontSize: 14,
    fontFamily: "Poppins_400Regular",
    color: "rgba(108, 108, 108, 0.89)",
  },
  informStatusContainer: {
    flex: 0.4,
  },
  textStatusContainer: {
    flex: 0.5,
    alignSelf: "center",
    paddingHorizontal: 10,
    justifyContent: "center",
    borderRadius: 10,
    borderWidth: 1.5,
  },
  textTimeContainer: {
    flex: 0.5,
    justifyContent: "center",
    alignItems: "center",
  },
  textStatus: {
    fontSize: 14,
    fontFamily: "Poppins_400Regular",
  },
  textTime: {
    fontSize: 12,
    fontFamily: "Poppins_400Regular",
    color: "rgba(108, 108, 108, 0.89)",
  },
  regular14: {
    fontSize: 14,
    color: "rgba(108, 108, 108, 0.89)",
    fontFamily: "Poppins_400Regular",
  },
  regular25In: {
    fontSize: 25,
    color: "rgba(52, 168, 83, 0.93)",
    fontFamily: "Poppins_400Regular",
    marginLeft: 10,
  },
  regular25Out: {
    fontSize: 25,
    color: "rgba(237, 74, 74, 0.87)",
    fontFamily: "Poppins_400Regular",
    marginLeft: 10,
  },
  regular14Add: {
    fontSize: 14,
    color: "rgba(27, 97, 181, 0.89)",
    fontFamily: "Poppins_400Regular",
    marginRight: 5,
  },
  modalContainer: {
    flex: 1,
    justifyContent: "center",
    alignItems: "center",
    backgroundColor: "rgba(0,0,0,0.5)",
  },
  modalContent: {
    padding: 16,
    borderRadius: 10,
    elevation: 5,
    width: 0.85 * screenWidth,
    backgroundColor: Colors.WHITE,
  },
  button: {
    marginTop: 5,
    paddingHorizontal: 10,
    paddingVertical: 4,
    borderRadius: 10,
    justifyContent: "center",
    backgroundColor: Colors.PRIMARY,
  },
  buttontext: {
    color: Colors.WHITE,
    fontFamily: "Poppins_400Regular",
    textAlign: "center",
    textAlignVertical: "center",
  },
  input: {
    padding: 4,
    marginTop: 10,
    marginBottom: 12,
    borderWidth: 1,
    borderColor: "rgba(27, 97, 181, 0.89)",
    borderRadius: 10,
  },
  label: {
    fontSize: 22,
    fontFamily: "Poppins_600SemiBold",
    textAlign: "center",
    marginBottom: 5,
  },
  content: {
    fontSize: 15,
    color: "rgba(27, 97, 181, 0.89)",
    fontFamily: "Poppins_600SemiBold",
    textAlign: "center",
    marginBottom: 5,
  },
  inform: {
    fontSize: 15,
    color: "rgba(108, 108, 108, 0.89)",
    fontFamily: "Poppins_400Regular",
    textAlign: "center",
    marginBottom: 5,
  },
  horizontalLine: {
    height: 1,
    width: "95%",
    alignSelf: "center",
    backgroundColor: "rgba(217, 217, 217, 0.89)",
    marginBottom: 10,
  },
});
