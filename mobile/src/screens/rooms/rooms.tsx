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
  TextInput,
  ViewStyle,
  StyleProp,
} from "react-native";
import { RoomsHeader } from "@/components/header/rooms-header";
import CalendarStrip from "react-native-calendar-strip";
import moment, { Moment } from "moment";
import Ionicons from "react-native-vector-icons/Ionicons";
import Octicons from "react-native-vector-icons/Octicons";
import { Colors } from "../../theme/variables";
import { rooms } from "@/services";
import axios from "axios";

const screenWidth = Dimensions.get("screen").width;
const screenHeight = Dimensions.get("screen").height;

interface Item {
  id: number;
  name: string;
  mssv: any;
  status: any;
  time: any;
}

interface InputPopupProps {
  isVisible: boolean;
  onClose: () => void;
  // check: any;
}

const InputPopup: React.FC<InputPopupProps> = ({
  isVisible,
  onClose,
  // check,
}) => {
  const [value, setValue] = useState("");
  const handleClick = () => {
    onClose();
    // check({ data: value });
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
            <Text style={styles.label}>Student ID:</Text>
            <TextInput
              placeholder="Type Student ID to check in/out"
              placeholderTextColor={"rgba(0, 0, 0, 0.5)"}
              value={value}
              onChangeText={setValue}
              style={styles.input}
              keyboardType="numeric"
            />
            <TouchableOpacity onPress={handleClick} style={styles.button}>
              <Text style={styles.buttontext}>OK</Text>
            </TouchableOpacity>
          </View>
        </View>
      </TouchableWithoutFeedback>
    </Modal>
  );
};


// async function fetchData() { 
//   try {
//     const roomData = await rooms.getroominform();

//   } catch (error) {
//     console.error('Error fetching room data:', error);
//   }
// }

export const Rooms = () => {
  const [selectedDate, setSelectedDate] = useState<Moment>(moment());
  const [isModalVisible, setIsModalVisible] = useState(false);

  const handleDateSelected = (date: Moment) => {
    setSelectedDate(date);
  };

  const handleAddRecordPress = () => {
    setIsModalVisible(true);
  };
  
  // useEffect(() => {
  //   fetchData();
  // }, []);

  const events: Item[] = [
    {
      id: 1,
      name: "Nguyen Van A",
      mssv: "2012345",
      status: "Check in",
      time: "06:30 am",
    },
    {
      id: 2,
      name: "Nguyen Van B",
      mssv: "2012989",
      status: "Check in",
      time: "06:30 pm",
    },
  ];

  const renderItem = ({ item }: { item: Item }) => (
    <TouchableOpacity style={styles.informContainer}>
      <View style={styles.informTextContainer}>
        <View style={styles.textNameContainer}>
          <Text style={styles.textName}>{item.name}</Text>
        </View>
        <View style={styles.textIdContainer}>
          <Text style={styles.textId}>MSSV: {item.mssv}</Text>
        </View>
      </View>
      <View style={styles.informStatusContainer}>
        <View style={styles.textStatusContainer}>
          <Text style={styles.textStatus}>{item.status}</Text>
        </View>
        <View style={styles.textTimeContainer}>
          <Text style={styles.textTime}>{item.time}</Text>
        </View>
      </View>
    </TouchableOpacity>
  );

  return (
    <View style={styles.container}>
      {/* <View style={styles.header_container}> */}
      <RoomsHeader />
      {/* </View> */}
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
                <Text style={styles.regular25In}>25</Text>
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
                <Text style={styles.regular25Out}>18</Text>
              </View>
              <Text style={styles.regular14}>check out</Text>
            </View>
          </View>
          <TouchableOpacity
            style={styles.addRecordStyle}
            onPress={handleAddRecordPress}
          >
            <Text style={styles.regular14Add}>Add record</Text>
            <Ionicons
              name="add-circle-outline"
              size={25}
              color="rgba(27, 97, 181, 0.89)"
            />
          </TouchableOpacity>
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
          // check={/* pass the check function here */}
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
  header_container: {
    // shadowOffset: {
    //   width: 0,
    //   height: 0,
    // },
    // shadowOpacity: 0.2,
    // shadowRadius: 5,
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
    borderTopLeftRadius: 10,
    borderTopRightRadius: 10,
    justifyContent: "center",
    backgroundColor: Colors.WHITE,
    shadowOffset: {
      width: 0,
      height: 0,
    },
    shadowOpacity: 0.1,
    shadowRadius: 2.5,
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
    marginVertical: 10,
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
    fontSize: 25,
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
    borderColor: "rgba(52, 168, 83, 0.89)",
    backgroundColor: "rgba(52, 168, 83, 0.15)",
  },
  textTimeContainer: {
    flex: 0.5,
    justifyContent: "center",
    alignItems: "center",
  },
  textStatus: {
    fontSize: 14,
    fontFamily: "Poppins_400Regular",
    color: "rgba(52, 168, 83, 0.93)",
  },
  textTime: {
    fontSize: 10,
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
    width: 0.8 * screenWidth,
    backgroundColor: Colors.WHITE,
  },
  button: {
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
    fontFamily: "Poppins_600SemiBold",
  },
});
