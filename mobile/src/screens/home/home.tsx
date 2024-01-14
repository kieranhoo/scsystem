import axios from "axios";
import React, { useEffect, useState } from "react";
import { View, Text, StyleSheet, Dimensions, FlatList, ScrollView , BackHandler } from "react-native";
import { BarChart } from 'react-native-gifted-charts';

import { Header } from "@/components/header";
import { Colors } from "@/theme/variables";


const screenWidth = Dimensions.get("screen").width;
const screenHeight = Dimensions.get("screen").height;

export const Home = () => {
  useEffect(() => {
    const backHandler = BackHandler.addEventListener(
      "hardwareBackPress",
      () => {
        return true;
      }
    );
    return () => backHandler.remove();
  }, []);
  const [bardata, setBardata] = useState({
    "month": "December",
    "year": "2023",
    "room_name": "sample",
    "room_id": "0",
    "data": [
      {
        "in": {
          "value": 1,
          "label": "5"
        },
        "out": {
          "value": 1,
          "label": "5"
        }
      }
    ]
  });

  const transformedArray = bardata.data.flatMap(item => ([
    {
      value: item.in.value,
      label: item.in.label,
      spacing: 2,
      labelWidth: 30,
      labelTextStyle: { color: 'gray' },
      frontColor: '#34A853',
      topLabelComponent: () => (
        <Text style={{ color: 'gray', fontSize: 10, marginBottom: 2 }}>
          {item.in.value}
        </Text>
      ),
    },
    {
      value: item.out.value,
      frontColor: '#ED4A4A',
      topLabelComponent: () => (
        <Text style={{ color: 'gray', fontSize: 10, marginBottom: 2 }}>
          {item.out.value}
        </Text>
      ),
    }
  ]));

  const DATA = [
    { room_id: '1', room_name: 'loading', in: 0, out: 0, total: 0 },
    { room_id: '2', room_name: 'loading', in: 0, out: 0, total: 0 },
    { room_id: '3', room_name: 'loading', in: 0, out: 0, total: 0 },
    // Add more items as needed
  ];

  const [generaldata, setGeneraldata] = useState(DATA);
  useEffect(() => {
    setTimeout(async () => {
      axios.get(`${process.env.BASE_URL}/stat/chart`, { params: { room_id: 1 } }).then(resp => {
        setBardata(resp.data);
      });
      axios.get(`${process.env.BASE_URL}/stat/rooms`).then((resp: any) => {
        setGeneraldata(resp.data.data);
      });
    }, generaldata == DATA ? 0 : 60000);
  }, []);

  type RoomProps = { room_name: string, in: number, out: number, total: number };

  const Room = (Room: RoomProps) => (
    <View style={styles.room}>
      <ScrollView horizontal>
        <Text style={styles.title}>{Room.room_name}</Text>
      </ScrollView>
      <View style={{ flexDirection: 'column', alignItems: 'center' }}>
        <View style={styles.total_card}>
          <Text style={{ fontWeight: 'bold', color: '#306fbb' }}>Total: {Room.total}</Text>
        </View >
        <View style={styles.in_card}>
          <Text style={{ fontWeight: 'bold', color: '#47b063' }}>In: {Room.in}</Text>
        </View>
        <View style={styles.out_card}>
          <Text style={{ fontWeight: 'bold', color: '#ef5e5e' }}>Out: {Room.out}</Text>
        </View>
      </View>
    </View>
  );

  const renderTitle = () => {
    return (
      <View style={styles.barchart_title}>
        <View style={{ flexDirection: 'row', alignItems: 'center' }}>
          <Text style={styles.title}>
            Daily in {bardata.room_name}
          </Text>
        </View>
        <View style={{ flexDirection: 'row' }}>
          <View style={{ flexDirection: 'row', alignItems: 'center' }}>
            <View style={styles.in_status} />
            <Text style={styles.text}>
              In
            </Text>
          </View>
          <View style={{ flexDirection: 'row', alignItems: 'center' }}>
            <View style={styles.out_status} />
            <Text style={styles.text}>
              Out
            </Text>
          </View>
        </View>
      </View>
    )
  }

  return (
    <View style={styles.container}>
      <View style={styles.header_container}>
        <Header title="Home" />
      </View>
      <View style={styles.content_container}>
        <View style={styles.component_content}>
          {renderTitle()}
          <BarChart
            data={transformedArray}
            height={150}
            barWidth={14}
            spacing={14}
            roundedTop
            roundedBottom
            hideRules
            xAxisThickness={0}
            yAxisThickness={-7}
            yAxisTextStyle={{ color: 'gray' }}
            noOfSections={5}
            maxValue={75}
          />
        </View>
        <View style={styles.component_content}>
          <FlatList
            data={generaldata}
            horizontal
            renderItem={({ item }) => <Room room_name={item.room_name} in={item.in} out={item.out} total={item.total} />}
            keyExtractor={item => item.room_id}
          />
        </View>
      </View>
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#f4f7fb',
  },
  header_container: {
    backgroundColor: Colors.WHITE,
    shadowOffset: {
      width: 0,
      height: 0,
    },
    shadowOpacity: 0.2,
    shadowRadius: 5,
  },
  barchart_title: {
    marginBottom: 20,
    flex: 1,
    flexDirection: 'row',
    justifyContent: 'space-between',
  },
  in_status: {
    height: 12,
    width: 12,
    borderRadius: 6,
    backgroundColor: '#34A853',
    marginRight: 8,
    marginBottom: 10,
  },
  out_status: {
    height: 12,
    width: 12,
    borderRadius: 6,
    backgroundColor: '#ED4A4A',
    marginRight: 8,
    marginBottom: 10,
  },
  text: {
    width: 30,
    height: 20,
    color: '#000000',
    marginBottom: 10,
  },
  content_container: {
    paddingTop: 30,
    paddingHorizontal: 0.05 * screenWidth,
  },
  component_content: {
    backgroundColor: 'white',
    padding: 20,
    paddingTop: 30,
    borderRadius: 10,
    marginBottom: 20,
    shadowOffset: {
      width: 0,
      height: 0,
    },
    shadowOpacity: 0.2,
    shadowRadius: 5,
  },
  room: {
    backgroundColor: Colors.WHITE,
    alignItems: 'center',
    padding: 10,
    width: 98,
    marginRight: 5,
    borderRadius: 10,
    borderWidth: 1,
    borderColor: '#7c7c7c',
  },
  total_card: {
    backgroundColor: '#d9e4f3',
    paddingTop: 10,
    paddingBottom: 10,
    alignItems: 'center',
    width: 80,
    borderTopStartRadius: 10,
    borderTopRightRadius: 10,
  },
  in_card: {
    width: 80,
    backgroundColor: '#e0f2e5',
    paddingTop: 10,
    paddingBottom: 10,
    alignItems: 'center',
  },
  out_card: {
    backgroundColor: '#fde4e4',
    paddingTop: 10,
    paddingBottom: 10,
    alignItems: 'center',
    width: 80,
    borderBottomLeftRadius: 10,
    borderBottomRightRadius: 10,
  },
  title: {
    fontSize: 18,
    height: 26,
    fontWeight: 'bold',
    color: '#000000',
    marginBottom: 10,
  },
});
