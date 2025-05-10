export interface Product {
  id: number;
  name: string;
  price: number;
  categoryID: number;
}

export interface ProductResponse {
  ID: number;
  Name: string;
  Price: number;
  CategoryID: number;
  Category:{
    ID: number;
    Name: string;
  }
}

export interface OrderItem {
  productId: number;
  quantity: number;
}

export interface OrderResponse {
  totalPrice: number;
  orderItems: {
    productId: number;
    quantity: number;
  }[];
}