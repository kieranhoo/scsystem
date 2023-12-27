import React, { useState } from "react";
import {
  View,
  Text,
  StyleSheet,
  Dimensions,
  TouchableOpacity,
  FlatList,
  ViewStyle,
  StyleProp
} from "react-native";
import { Header } from "@/components/header";
import CalendarStrip from "react-native-calendar-strip";
import moment, { Moment } from "moment";
import Ionicons from "react-native-vector-icons/Ionicons";
import Octicons from "react-native-vector-icons/Octicons";
import { Colors } from "../../theme/variables";
const screenWidth = Dimensions.get("screen").width;
const screenHeight = Dimensions.get("screen").height;

interface Item {
  id: number;
  name: string;
  mssv: any;
  status: any;
  time: any;
}

export const Rooms = () => {
  const [selectedDate, setSelectedDate] = useState<Moment>(moment());

  const handleDateSelected = (date: Moment) => {
    setSelectedDate(date);
  };

  const events: Item[] = [
    {
      id: 1,
      name: 'Nguyen Van A',
      mssv: 'MSSV: 2012345',
      status: "checkin",
      time: "06:30 am",
    },
    {
      id: 2,
      name: 'Nguyen Van B',
      mssv: 'MSSV: 2012989',
      status: "checkin",
      time: "06:30 pm",
    },
  ];

  const renderItem = ({ item }: { item: Item }) => (
    <View>
      <Text>{item.name}</Text>
    </View>
  );

  return (
    <View style={styles.container}>
      <View style={styles.header_container}>
        <Header title="Rooms" />
      </View>
      <View style={styles.content_container}>
        <View style={styles.calendarContainer}>
          <CalendarStrip
            style={styles.calendarStripStyle}
            showMonth={false}
            scrollable
            scrollerPaging
            numDaysInWeek={7}

            dateNumberStyle={{color: "#6C6C6C"}}
            dateNameStyle={{color: "#6C6C6C"}}
            dayContainerStyle={styles.dateContainer}
            highlightDateContainerStyle={styles.highlightDateContainer}
            highlightDateNameStyle={{color: "#1B61B5"}}
            highlightDateNumberStyle={{color: "#1B61B5"}}
            selectedDate={selectedDate}
            // onDateSelected={handleDateSelected}
            // calendarAnimation={{ type: "sequence", duration: 30 }}
            // daySelectionAnimation={{
            //   type: 'border',
            //   duration: 20,
            //   borderWidth: 1.5,
            //   borderHighlightColor: "rgba(27, 97, 181, 0.89)",
            // }}
            // iconContainer={{flex: 0.1}}
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
          <TouchableOpacity style={styles.addRecordStyle}>
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
            // contentContainerStyle={styles.eventList}
          />
        </View>
      </View>
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
    borderRadius: 5,
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
    borderTopLeftRadius: 5,
    borderTopRightRadius: 5,
    justifyContent: "center",
    backgroundColor: Colors.WHITE,
    shadowOffset: {
      width: 0,
      height: -3,
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
    borderBottomLeftRadius: 5,
    borderBottomRightRadius: 5,
    backgroundColor: Colors.WHITE,
    shadowOffset: {
      width: 0,
      height: 3,
    },
    shadowOpacity: 0.2,
    shadowRadius: 5,
  },
  detailContainer: {
    flex: 0.69,
    backgroundColor: "yellow",
    marginTop: 10,
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
});
